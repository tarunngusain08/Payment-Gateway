package middleware

import (
	"io"
	"log"
	"net/http"
	"regexp"
	"strings"
	"time"
)

// maskPII masks card numbers and user info in the input string.
func maskPII(input string) string {
	// Mask card numbers (simple regex for 16-digit numbers)
	cardRegex := regexp.MustCompile(`\b(\d{4})\d{8,10}(\d{4})\b`)
	input = cardRegex.ReplaceAllString(input, "$1********$2")
	// Mask email addresses
	emailRegex := regexp.MustCompile(`([a-zA-Z0-9._%+-]+)@([a-zA-Z0-9.-]+\.[a-zA-Z]{2,})`)
	input = emailRegex.ReplaceAllString(input, "****@$2")
	return input
}

// LoggingMiddleware logs requests with PII masking.
func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		var bodyCopy string
		if r.Body != nil {
			bodyBytes, _ := io.ReadAll(r.Body)
			bodyCopy = string(bodyBytes)
			r.Body = io.NopCloser(strings.NewReader(bodyCopy))
		}
		log.Printf("Request: %s %s Body: %s", r.Method, r.URL.Path, maskPII(bodyCopy))
		next.ServeHTTP(w, r)
		log.Printf("Completed %s %s in %v", r.Method, r.URL.Path, time.Since(start))
	})
}
