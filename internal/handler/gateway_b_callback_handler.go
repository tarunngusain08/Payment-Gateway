package handler

import (
	"Payment-Gateway/internal/dtos"
	"Payment-Gateway/internal/service"
	"encoding/xml"
	"io"
	"net/http"
)

type GatewayBCallbackHandler struct {
	Service service.Callback
}

func NewGatewayBCallback(service service.Callback) GatewayBCallbackHandler {
	return GatewayBCallbackHandler{Service: service}
}

func (h *GatewayBCallbackHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "invalid XML", http.StatusBadRequest)
		return
	}
	var req dtos.HandleCallbackRequest
	if err := xml.Unmarshal(body, &req); err != nil {
		http.Error(w, "invalid XML", http.StatusBadRequest)
		return
	}
	resp, err := h.Service.HandleCallback(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	xml.NewEncoder(w).Encode(resp)
}
