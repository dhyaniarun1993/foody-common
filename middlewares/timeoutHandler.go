package middlewares

import (
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

// TimeoutHandler wraps http.Handler and handle request timeout
func TimeoutHandler(timeout time.Duration) mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		handlerFunc := func(w http.ResponseWriter, r *http.Request) {
			var timeoutHandler http.Handler
			msg := "Your request has timed out."
			timeoutHandler = http.TimeoutHandler(next, timeout, msg)
			timeoutHandler.ServeHTTP(w, r)
		}
		return http.HandlerFunc(handlerFunc)
	}
}
