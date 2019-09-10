package middlewares

import (
	"net/http"

	"github.com/gorilla/mux"
)

func convertToHandler(handlerFunc http.HandlerFunc) http.Handler {
	return http.HandlerFunc(handlerFunc)
}

// ChainHandlerFuncMiddlewares chains handlerFunc and middlewares
func ChainHandlerFuncMiddlewares(handlerFunc http.HandlerFunc, middlewares ...mux.MiddlewareFunc) http.Handler {
	handler := convertToHandler(handlerFunc)
	for i := len(middlewares) - 1; i >= 0; i-- {
		handler = middlewares[i](handler)
	}
	return handler
}

// ChainHandlerMiddlewares chains handlerFunc and middlewares
func ChainHandlerMiddlewares(handler http.Handler, middlewares ...mux.MiddlewareFunc) http.Handler {
	for i := len(middlewares) - 1; i >= 0; i-- {
		handler = middlewares[i](handler)
	}
	return handler
}
