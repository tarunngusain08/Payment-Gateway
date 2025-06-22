package gateway

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestGatewayA_ProcessDeposit_Success(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]interface{}{"result": "ok"})
	}))
	defer ts.Close()

	g := NewGatewayA(ts.URL, "gatewayA")
	resp, err := g.ProcessDeposit(nil)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	m, ok := resp.(map[string]interface{})
	if !ok || m["result"] != "ok" {
		t.Errorf("unexpected response: %v", resp)
	}
}

func TestGatewayA_ProcessDeposit_Failure(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusBadRequest)
	}))
	defer ts.Close()

	g := NewGatewayA(ts.URL, "gatewayA")
	_, err := g.ProcessDeposit(nil)
	if err == nil || err.Error() != "gateway A failure" {
		t.Errorf("expected gateway A failure error, got %v", err)
	}
}

func TestGatewayA_ProcessDeposit_Timeout(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(3 * time.Second)
	}))
	defer ts.Close()

	g := NewGatewayA(ts.URL, "gatewayA")
	_, err := g.ProcessDeposit(nil)
	if err == nil || err.Error() != "gateway A timeout" {
		t.Errorf("expected gateway A timeout error, got %v", err)
	}
}

func TestGatewayA_ProcessWithdrawal_Success(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]interface{}{"result": "ok"})
	}))
	defer ts.Close()

	g := NewGatewayA(ts.URL, "gatewayA")
	resp, err := g.ProcessWithdrawal(nil)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	m, ok := resp.(map[string]interface{})
	if !ok || m["result"] != "ok" {
		t.Errorf("unexpected response: %v", resp)
	}
}

func TestGatewayA_ProcessWithdrawal_Failure(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusBadRequest)
	}))
	defer ts.Close()

	g := NewGatewayA(ts.URL, "gatewayA")
	_, err := g.ProcessWithdrawal(nil)
	if err == nil || err.Error() != "gateway A failure" {
		t.Errorf("expected gateway A failure error, got %v", err)
	}
}

func TestGatewayA_ProcessWithdrawal_Timeout(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(3 * time.Second)
	}))
	defer ts.Close()

	g := NewGatewayA(ts.URL, "gatewayA")
	_, err := g.ProcessWithdrawal(nil)
	if err == nil || err.Error() != "gateway A timeout" {
		t.Errorf("expected gateway A timeout error, got %v", err)
	}
}
