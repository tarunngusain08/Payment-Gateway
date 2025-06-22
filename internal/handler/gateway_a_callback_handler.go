package handler

import (
	"Payment-Gateway/internal/dtos"
	"Payment-Gateway/internal/middleware"
	"Payment-Gateway/internal/service"
	"encoding/json"
	"net/http"

	"go.uber.org/zap"
)

type GatewayACallbackHandler struct {
	Service service.Callback
}

func NewGatewayACallback(service service.Callback) GatewayACallbackHandler {
	return GatewayACallbackHandler{Service: service}
}

func (h *GatewayACallbackHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log := middleware.LoggerFromContext(r.Context()).With(zap.String("func", "GatewayACallbackHandler.ServeHTTP"))
	log.Info("Received GatewayA callback")

	var req dtos.HandleCallbackRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		log.Warn("Invalid GatewayA callback JSON", zap.Error(err))
		http.Error(w, "invalid JSON", http.StatusBadRequest)
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
		log.Warn("Validation failed for GatewayA callback", zap.Error(err))
		http.Error(w, "Validation failed", http.StatusBadRequest)
		return
	}

	resp, err := h.Service.HandleCallback(req)
	if err != nil {
		log.Error("GatewayA callback processing failed", zap.Error(err))
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	log.Info("GatewayA callback processed successfully")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}
