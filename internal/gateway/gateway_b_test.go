package gateway

import (
	"Payment-Gateway/internal/dtos"
	"encoding/xml"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestGatewayB_ProcessDeposit_Success(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		resp := dtos.SOAPEnvelope{
			Body: dtos.SOAPBody{
				DepositResponse: &dtos.SOAPDepositResponse{Result: "ok"},
			},
		}
		xml.NewEncoder(w).Encode(resp)
	}))
	defer ts.Close()

	g := NewGatewayB(ts.URL, "gatewayB", getTestResilienceConfig())
	resp, err := g.ProcessDeposit(nil)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	_, ok := resp.(dtos.SOAPEnvelope)
	if !ok {
		t.Errorf("unexpected response type: %T", resp)
	}
}

func TestGatewayB_ProcessDeposit_Failure(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusBadRequest)
	}))
	defer ts.Close()

	g := NewGatewayB(ts.URL, "gatewayB", getTestResilienceConfig())
	_, err := g.ProcessDeposit(nil)
	if err == nil || err.Error() != "gateway B failure" {
		t.Errorf("expected gateway B failure error, got %v", err)
	}
}

func TestGatewayB_ProcessDeposit_Timeout(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(3 * time.Second)
	}))
	defer ts.Close()

	g := NewGatewayB(ts.URL, "gatewayB", getTestResilienceConfig())
	_, err := g.ProcessDeposit(nil)
	if err == nil || err.Error() != "gateway B timeout" {
		t.Errorf("expected gateway B timeout error, got %v", err)
	}
}

func TestGatewayB_ProcessWithdrawal_Success(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		resp := dtos.SOAPEnvelope{
			Body: dtos.SOAPBody{
				WithdrawalResponse: &dtos.SOAPWithdrawalResponse{Result: "ok"},
			},
		}
		xml.NewEncoder(w).Encode(resp)
	}))
	defer ts.Close()

	g := NewGatewayB(ts.URL, "gatewayB", getTestResilienceConfig())
	resp, err := g.ProcessWithdrawal(nil)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	_, ok := resp.(dtos.SOAPEnvelope)
	if !ok {
		t.Errorf("unexpected response type: %T", resp)
	}
}

func TestGatewayB_ProcessWithdrawal_Failure(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusBadRequest)
	}))
	defer ts.Close()

	g := NewGatewayB(ts.URL, "gatewayB", getTestResilienceConfig())
	_, err := g.ProcessWithdrawal(nil)
	if err == nil || err.Error() != "gateway B failure" {
		t.Errorf("expected gateway B failure error, got %v", err)
	}
}

func TestGatewayB_ProcessWithdrawal_Timeout(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(3 * time.Second)
	}))
	defer ts.Close()

	g := NewGatewayB(ts.URL, "gatewayB", getTestResilienceConfig())
	_, err := g.ProcessWithdrawal(nil)
	if err == nil || err.Error() != "gateway B timeout" {
		t.Errorf("expected gateway B timeout error, got %v", err)
	}
}
