package mwLogger

import (
	"net/http"
	"os"
	"time"

	"github.com/charmbracelet/log"
	"github.com/go-chi/chi/v5/middleware"
)

func New(logger *log.Logger) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {

		logger = log.NewWithOptions(os.Stdout, log.Options{
			Prefix: " component=middleware/logger",
		})

		logger.Info("Logger middleware enabled")

		fn := func(w http.ResponseWriter, r *http.Request) {
			entry := logger.With(
				"method", r.Method,
				"path", r.URL.Path,
				"remote_addr", r.RemoteAddr,
				"user_agent", r.UserAgent(),
				"request_id", middleware.GetReqID(r.Context()),
			)
			ww := middleware.NewWrapResponseWriter(w, r.ProtoMajor)

			t1 := time.Now()
			defer func() {
				duration := time.Since(t1)
				entry.With(
					"status", ww.Status(),
					"bytes", ww.BytesWritten(),
					"duration", duration.String(),
				).Info("Request completed")
			}()

			next.ServeHTTP(ww, r)
		}
		return http.HandlerFunc(fn)
	}
}
