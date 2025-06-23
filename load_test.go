package paymentgateway

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"math/rand"
	"net/http"
	"sync"
	"testing"
	"time"
)

// --- Load Test Parameters ---
const (
	baseURL             = "http://localhost:8000"
	numTransactions     = 1000
	concurrentClients   = 10
	callbackDelayMillis = 100 // max random delay in milliseconds
)

// --- Transaction & Callback Payloads ---
type TransactionRequest struct {
	AccountID string  `json:"account_id"`
	Amount    float64 `json:"amount"`
}

type CallbackRequest struct {
	TransactionID string  `json:"transactionId" xml:"TransactionID"`
	GatewayRef    string  `json:"gatewayRef" xml:"GatewayRef"`
	Amount        float64 `json:"amount" xml:"Amount"`
	Currency      string  `json:"currency" xml:"Currency"`
	Status        string  `json:"status" xml:"Status"`
}

// --- Helper Functions ---
func sendTransaction(t *testing.T) (string, error) {
	payload := TransactionRequest{
		AccountID: fmt.Sprintf("acc-%d", rand.Intn(1000)),
		Amount:    rand.Float64() * 100,
	}
	body, _ := json.Marshal(payload)
	resp, err := http.Post(fmt.Sprintf("%s/deposit", baseURL), "application/json", bytes.NewReader(body))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	var result struct {
		Success bool   `json:"success"`
		Message string `json:"message,omitempty"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return "", err
	}
	if !result.Success {
		return "", fmt.Errorf("deposit failed: %s", result.Message)
	}
	// Return a fake transaction ID for callback simulation
	return fmt.Sprintf("tx-%d", rand.Intn(1000000)), nil
}

func sendCallback(t *testing.T, gateway string, transactionID string) {
	payload := CallbackRequest{
		TransactionID: transactionID,
		GatewayRef:    fmt.Sprintf("ref-%d", rand.Intn(10000)),
		Amount:        rand.Float64() * 100,
		Currency:      "USD",
		Status:        "SUCCESS",
	}
	if gateway == "gatewayA" {
		body, _ := json.Marshal(payload)
		url := fmt.Sprintf("%s/callback/gateway-a", baseURL)
		http.Post(url, "application/json", bytes.NewReader(body))
	} else {
		body, _ := xml.Marshal(payload)
		url := fmt.Sprintf("%s/callback/gateway-b", baseURL)
		http.Post(url, "application/xml", bytes.NewReader(body))
	}
}

// --- Load Test ---
func TestLoadSimulator(t *testing.T) {
	var wg sync.WaitGroup
	sem := make(chan struct{}, concurrentClients)
	for i := 0; i < numTransactions; i++ {
		wg.Add(1)
		sem <- struct{}{}
		go func() {
			defer wg.Done()
			defer func() { <-sem }()
			transactionID, err := sendTransaction(t)
			if err != nil || transactionID == "" {
				t.Logf("Transaction failed: %v", err)
				return
			}
			// Randomly trigger callback for either gatewayA or gatewayB with delay between 1 and callbackDelayMillis milliseconds
			go func(txID string) {
				delay := time.Duration(rand.Intn(callbackDelayMillis-1)+1) * time.Millisecond
				time.Sleep(delay)
				gateway := "gatewayA"
				if rand.Intn(2) == 0 {
					gateway = "gatewayB"
				}
				sendCallback(t, gateway, txID)
			}(transactionID)
		}()
	}
	wg.Wait()
}
