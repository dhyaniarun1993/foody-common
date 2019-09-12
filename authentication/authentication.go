package authentication

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type key string

const (
	userIDHeader   = "X-User-ID"
	userRoleHeader = "X-User-Role"
	appIDHeader    = "X-App-ID"
	userIDKey      = key("userId")
	userRoleKey    = key("userRole")
	appIDKey       = key("appId")
)

// AuthHandler wraps http.Handler and handle
func AuthHandler() mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		handlerFunc := func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			userID := r.Header.Get(userIDHeader)
			userRole := r.Header.Get(userRoleHeader)
			appID := r.Header.Get(appIDHeader)
			if userID == "" || appID == "" {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusUnauthorized)
				fmt.Fprintf(w, `{"message": %q}`, "Auth info missing.")
				return
			}

			ctx = context.WithValue(ctx, userIDKey, userID)
			ctx = context.WithValue(ctx, userRoleKey, userRole)
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

// GetUserRole extracts and return user role from context
func GetUserRole(ctx context.Context) (string, bool) {
	id, ok := ctx.Value(userRoleKey).(string)
	return id, ok
}

// GetAppID extracts and return app id from context
func GetAppID(ctx context.Context) (string, bool) {
	id, ok := ctx.Value(appIDKey).(string)
	return id, ok
}
