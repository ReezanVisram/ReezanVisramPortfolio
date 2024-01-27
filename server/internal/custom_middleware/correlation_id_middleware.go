package custom_middleware

import (
	"context"
	"net/http"
	"reezanvisramportfolio/internal/custom_logging"

	"github.com/google/uuid"
)

func CorrelationIdMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		ctx := context.WithValue(r.Context(), custom_logging.KeyCorrelationId, uuid.New().String())

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
