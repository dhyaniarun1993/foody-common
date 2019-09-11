package middlewares

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type key string

const (
	userIDKey = key("userId")
	appIDKey  = key("appId")
)

// Authentication provides http.Handler for authentication
type Authentication struct {
	UserIDHeader string
	AppIDHeader  string
}

// AuthHandler wraps http.Handler and handle
func (config Authentication) AuthHandler() mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		handlerFunc := func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			userID := r.Header.Get(config.UserIDHeader)
			appID := r.Header.Get(config.AppIDHeader)
			if userID == "" || appID == "" {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusUnauthorized)
				fmt.Fprintf(w, `{"message": %q}`, "Auth info missing.")
				return
			}

			ctx = context.WithValue(ctx, userIDKey, userID)
			ctx = context.WithValue(ctx, appIDKey, appID)
			r = r.WithContext(ctx)
			next.ServeHTTP(w, r)
		}
		return http.HandlerFunc(handlerFunc)
	}
}

// GetUserID extracts and return user id from context
func GetUserID(ctx context.Context) (string, bool) {
	id, ok := ctx.Value(userIDKey).(string)
	return id, ok
}

// GetAppID extracts and return app id from context
func GetAppID(ctx context.Context) (string, bool) {
	id, ok := ctx.Value(appIDKey).(string)
	return id, ok
}
