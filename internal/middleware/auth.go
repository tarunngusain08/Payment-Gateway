package middleware

import "net/http"

// AuthMiddleware is a stub for authentication logic.
func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Example: Check for a static API key header (for demonstration)
		if r.Header.Get("X-API-Key") == "" {
			http.Error(w, "unauthorized", http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	})
}
