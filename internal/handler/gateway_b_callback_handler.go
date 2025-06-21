package handler

import (
	"Payment-Gateway/internal/dtos"
	"Payment-Gateway/internal/service"
	"encoding/xml"
	"io"
	"net/http"
)

type GatewayBCallbackHandler struct {
	Service service.GatewayBCallbackService
}

func NewGatewayBCallback(service service.GatewayBCallbackService) GatewayACallbackHandler {
	return GatewayACallbackHandler{Service: service}
}

func (h *GatewayBCallbackHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "invalid XML", http.StatusBadRequest)
		return
	}
	var req dtos.GatewayBCallbackRequest
	if err := xml.Unmarshal(body, &req); err != nil {
		http.Error(w, "invalid XML", http.StatusBadRequest)
		return
	}
	resp, status := h.Service.HandleCallback(req)
	w.WriteHeader(status)
	xml.NewEncoder(w).Encode(resp)
}
