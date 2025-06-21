package middleware

import (
	"log"
	"net/http"
	"time"
)

// LatencyTrackerMiddleware tracks response time for each request.
func LatencyTrackerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		log.Printf("Latency for %s: %v", r.URL.Path, time.Since(start))
	})
}
