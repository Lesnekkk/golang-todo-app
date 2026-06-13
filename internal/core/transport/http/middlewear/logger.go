package core_http_middleware

import (
	"net/http"
	"time"

	"go.uber.org/zap"
)

func Logger(logger *zap.Logger) Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			start := time.Now()

			next.ServeHTTP(w, r)

			logger.Info("request",
				zap.String("method", r.Method),
				zap.String("path", r.URL.Path),
				zap.String("request_id", r.Header.Get("X-Request-ID")),
				zap.Duration("duration", time.Since(start)),
			)
		})
	}
}
