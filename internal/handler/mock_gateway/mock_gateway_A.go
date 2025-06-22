package mockgateway

import (
	"encoding/json"
	"net/http"
)

func GatewayAMockHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	resp := map[string]interface{}{
		"status":  "success",
		"message": "Mock Gateway A processed the request successfully",
	}
	json.NewEncoder(w).Encode(resp)
}

func GatewayAMockDepositHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	resp := map[string]interface{}{
		"status":  "success",
		"message": "Mock Gateway A processed the deposit successfully",
	}
	json.NewEncoder(w).Encode(resp)
}

func GatewayAMockWithdrawalHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	resp := map[string]interface{}{
		"status":  "success",
		"message": "Mock Gateway A processed the withdrawal successfully",
	}
	json.NewEncoder(w).Encode(resp)
}
