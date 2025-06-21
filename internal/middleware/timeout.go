package middleware

import (
	"net/http"
	"time"
)

// TimeoutMiddleware sets a timeout for each request.
func TimeoutMiddleware(timeout time.Duration) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.TimeoutHandler(next, timeout, "request timed out")
	}
}
