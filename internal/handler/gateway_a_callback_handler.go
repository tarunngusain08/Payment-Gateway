package handler

import (
	"Payment-Gateway/internal/dtos"
	"Payment-Gateway/internal/service"
	"encoding/json"
	"net/http"
)

type GatewayACallbackHandler struct {
	Service service.Callback
}

func NewGatewayACallback(service service.Callback) GatewayACallbackHandler {
	return GatewayACallbackHandler{Service: service}
}

func (h *GatewayACallbackHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var req dtos.HandleCallbackRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid JSON", http.StatusBadRequest)
		return
	}
	resp, err := h.Service.HandleCallback(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}
