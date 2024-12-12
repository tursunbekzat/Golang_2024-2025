package middlewares

import (
	"net/http"

	"golang.org/x/time/rate"
)

var rateLimiter = rate.NewLimiter(1, 5) // 1 request/second, burst size of 5

// RateLimitMiddleware limits the number of requests to an endpoint
func RateLimitMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !rateLimiter.Allow() {
			http.Error(w, "Too many requests", http.StatusTooManyRequests)
			return
		}
		next.ServeHTTP(w, r)
	})
}
