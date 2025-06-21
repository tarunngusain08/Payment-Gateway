package handler

import (
	"Payment-Gateway/internal/dtos"
	"Payment-Gateway/internal/service"
	"encoding/json"
	"net/http"
)

type GatewayACallbackHandler struct {
	Service service.GatewayACallbackService
}

func NewGatewayACallback(service service.GatewayACallbackService) GatewayACallbackHandler {
	return GatewayACallbackHandler{Service: service}
}

func (h *GatewayACallbackHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var req dtos.GatewayACallbackRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid JSON", http.StatusBadRequest)
		return
	}
	resp, status := h.Service.HandleCallback(req)
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(resp)
}
