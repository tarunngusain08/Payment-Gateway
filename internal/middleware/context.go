package middleware

import (
	"Payment-Gateway/pkg/logger"
	"context"
	"net/http"

	"github.com/google/uuid"
	"go.uber.org/zap"
)

type contextKey string

const (
	ContextKeyTraceID   contextKey = "trace-id"
	ContextKeyRequestID contextKey = "request-id"
	ContextKeyLogger    contextKey = "logger"
)

// ContextMiddleware injects trace-id, request-id, and logger into the request context.
func ContextMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		traceID := r.Header.Get("X-Trace-Id")
		if traceID == "" {
			traceID = uuid.NewString()
		}
		requestID := r.Header.Get("X-Request-Id")
		if requestID == "" {
			requestID = uuid.NewString()
		}

		ctx := context.WithValue(r.Context(), ContextKeyTraceID, traceID)
		ctx = context.WithValue(ctx, ContextKeyRequestID, requestID)

		// Attach a logger with trace-id and request-id fields
		reqLogger := logger.GetLogger().With(
			zap.String("trace_id", traceID),
			zap.String("request_id", requestID),
		)
		ctx = context.WithValue(ctx, ContextKeyLogger, reqLogger)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// LoggerFromContext extracts the zap.Logger from context, or returns the global logger if missing.
func LoggerFromContext(ctx context.Context) *zap.Logger {
	l, ok := ctx.Value(ContextKeyLogger).(*zap.Logger)
	if !ok || l == nil {
		return logger.GetLogger()
	}
	return l
}
