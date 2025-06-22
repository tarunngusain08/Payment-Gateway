package handler

import (
	"Payment-Gateway/internal/dtos"
	"Payment-Gateway/internal/middleware"
	"Payment-Gateway/internal/service"
	"encoding/xml"
	"io"
	"net/http"

	"go.uber.org/zap"
)

type GatewayBCallbackHandler struct {
	Service service.Callback
}

func NewGatewayBCallback(service service.Callback) GatewayBCallbackHandler {
	return GatewayBCallbackHandler{Service: service}
}

func (h *GatewayBCallbackHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log := middleware.LoggerFromContext(r.Context()).With(zap.String("func", "GatewayBCallbackHandler.ServeHTTP"))
	log.Info("Received GatewayB callback")

	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Warn("Failed to read GatewayB callback body", zap.Error(err))
		http.Error(w, "invalid XML", http.StatusBadRequest)
		return
	}
	var req dtos.HandleCallbackRequest
	if err := xml.Unmarshal(body, &req); err != nil {
		log.Warn("Invalid GatewayB callback XML", zap.Error(err))
		http.Error(w, "invalid XML", http.StatusBadRequest)
		return
	}

	log = log.With(
		zap.String("transaction_id", req.TransactionID),
		zap.String("gateway_ref", req.GatewayRef),
		zap.Float64("amount", req.Amount),
		zap.String("currency", req.Currency),
		zap.String("status", req.Status),
	)

	if err := req.Validate(); err != nil {
		log.Warn("Validation failed for GatewayB callback", zap.Error(err))
		http.Error(w, "Validation failed", http.StatusBadRequest)
		return
	}

	resp, err := h.Service.HandleCallback(req)
	if err != nil {
		log.Error("GatewayB callback processing failed", zap.Error(err))
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	log.Info("GatewayB callback processed successfully")
	w.WriteHeader(http.StatusOK)
	xml.NewEncoder(w).Encode(resp)
}
