package handler

import (
	"Payment-Gateway/internal/cache"
	"Payment-Gateway/internal/dtos"
	"Payment-Gateway/internal/middleware"
	"Payment-Gateway/internal/service"
	"encoding/xml"
	"fmt"
	"io"
	"net/http"

	"go.uber.org/zap"
)

type GatewayBCallbackHandler struct {
	Service service.Callback
	Cache   cache.CacheStore
}

func NewGatewayBCallback(service service.Callback, c cache.CacheStore) GatewayBCallbackHandler {
	return GatewayBCallbackHandler{
		Service: service,
		Cache:   c,
	}
}

func (h *GatewayBCallbackHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := middleware.LoggerFromContext(ctx).With(zap.String("func", "GatewayBCallbackHandler.ServeHTTP"))
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

	cacheKey := fmt.Sprintf("callback:gatewayB:%s:%s", req.TransactionID, req.GatewayRef)
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

	if cachedResp, found := h.Cache.Get(ctx, cacheKey); found {
		log.Info("Returning cached response for duplicate callback")
		w.WriteHeader(http.StatusOK)
		xml.NewEncoder(w).Encode(cachedResp)
		return
	}

	resp, err := h.Service.HandleCallback(req)
	if err != nil {
		log.Error("GatewayB callback processing failed", zap.Error(err))
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := h.Cache.Set(ctx, cacheKey, resp); err != nil {
		log.Error("Failed to cache callback response", zap.Error(err))
	}

	log.Info("GatewayB callback processed successfully")
	w.WriteHeader(http.StatusOK)
	xml.NewEncoder(w).Encode(resp)
}
