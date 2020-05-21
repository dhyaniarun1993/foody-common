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
	clientIDHeader = "X-Client-ID"
	authKey        = key("auth")
)

// AuthHandler wraps http.Handler and handle
func AuthHandler() mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		handlerFunc := func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			userID := r.Header.Get(UserIDHeader)
			userRole := r.Header.Get(UserRoleHeader)
			clientID := r.Header.Get(clientIDHeader)
			if userID == "" || clientID == "" || userRole == "" {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusUnauthorized)
				fmt.Fprintf(w, `{"message": %q}`, "Auth info missing.")
				return
			}

			auth := Auth{
				clientID: clientID,
				userID:   userID,
				userRole: userRole,
			}
			ctx = context.WithValue(ctx, authKey, auth)
			r = r.WithContext(ctx)
			next.ServeHTTP(w, r)
		}
		return http.HandlerFunc(handlerFunc)
	}
}

// Auth provides model definition for Auth
type Auth struct {
	clientID string
	userID   string
	userRole string
}

// GetAuthFromContext extracts and return Auth object from context
func GetAuthFromContext(ctx context.Context) (Auth, bool) {
	auth, ok := ctx.Value(authKey).(Auth)
	return auth, ok
}

// GetUserID returns the userID of the user
func (auth *Auth) GetUserID() string {
	return auth.userID
}

// GetUserRole returns the role of the user
func (auth *Auth) GetUserRole() string {
	return auth.userRole
}

// GetClientID returns the ClientID of the client that requested the resource
func (auth *Auth) GetClientID() string {
	return auth.clientID
}
