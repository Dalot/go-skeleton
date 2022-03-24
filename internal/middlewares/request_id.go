package middlewares

import (
	"context"
	"net/http"

	"github.com/dalot/go-skeleton/pkg/constants"
	"github.com/google/uuid"
)

type ReqIDContextKey string

var reqIDContextKey = ReqIDContextKey("request_id")

// requestIDContext creates a context with request id
func requestIDContext(ctx context.Context, rid string) context.Context {
	return context.WithValue(ctx, reqIDContextKey, rid)
}

// requestIDFromContext returns the request id from context
func requestIDFromContext(ctx context.Context) string {
	rid, ok := ctx.Value(reqIDContextKey).(string)
	if !ok {
		return ""
	}
	return rid
}

// RequestIDHandler sets unique request id.
// If header `X-Request-ID` is already present in the request, that is considered the
// request id. Otherwise, generates a new unique ID.
func RequestIDHandler(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		rid := r.Header.Get(constants.HeaderKeyRequestID)

		if rid == "" {
			rid = uuid.New().String()
		}

		w.Header().Set(constants.HeaderKeyRequestID, rid)

		ctx := requestIDContext(r.Context(), rid)
		h.ServeHTTP(w, r.WithContext(ctx))
	})
}
