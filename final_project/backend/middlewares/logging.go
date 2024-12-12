package middlewares

import (
	"net/http"
	"time"

	log "github.com/sirupsen/logrus"
)

func RequestLoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		duration := time.Since(start)

		log.WithFields(log.Fields{
			"method":     r.Method,
			"path":       r.URL.Path,
			"client_ip":  r.RemoteAddr,
			"duration":   duration.Seconds(),
			"user_agent": r.UserAgent(),
		}).Info("Request completed")
	})
}
