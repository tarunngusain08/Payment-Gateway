package middleware

import (
	"context"
	"net/http"

	"github.com/google/uuid"
)

type contextKey string

const (
	ContextKeyTraceID   contextKey = "trace-id"
	ContextKeyRequestID contextKey = "request-id"
)

// ContextMiddleware injects trace-id and request-id into the request context.
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
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
