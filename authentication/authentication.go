package authentication

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type key string

// authentication constants
const (
	UserIDHeader   = "X-User-ID"
	UserRoleHeader = "X-User-Role"
	AppIDHeader    = "X-App-ID"
	UserIDKey      = key("userId")
	UserRoleKey    = key("userRole")
	AppIDKey       = key("appId")
)

// AuthHandler wraps http.Handler and handle
func AuthHandler() mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		handlerFunc := func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			userID := r.Header.Get(UserIDHeader)
			userRole := r.Header.Get(UserRoleHeader)
			appID := r.Header.Get(AppIDHeader)
			if userID == "" || appID == "" || userRole == "" {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusUnauthorized)
				fmt.Fprintf(w, `{"message": %q}`, "Auth info missing.")
				return
			}

			ctx = context.WithValue(ctx, UserIDKey, userID)
			ctx = context.WithValue(ctx, UserRoleKey, userRole)
			ctx = context.WithValue(ctx, AppIDKey, appID)
			r = r.WithContext(ctx)
			next.ServeHTTP(w, r)
		}
		return http.HandlerFunc(handlerFunc)
	}
}

// GetUserID extracts and return user id from context
func GetUserID(ctx context.Context) (string, bool) {
	id, ok := ctx.Value(UserIDKey).(string)
	return id, ok
}

// GetUserRole extracts and return user role from context
func GetUserRole(ctx context.Context) (string, bool) {
	id, ok := ctx.Value(UserRoleKey).(string)
	return id, ok
}

// GetAppID extracts and return app id from context
func GetAppID(ctx context.Context) (string, bool) {
	id, ok := ctx.Value(AppIDKey).(string)
	return id, ok
}
