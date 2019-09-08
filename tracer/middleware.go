package tracer

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	"github.com/uber/jaeger-client-go"
)

const (
	defaultComponentName     = "net/http"
	defaultTracingHTTPHeader = "x-request-id"
)

type statusRecorder struct {
	http.ResponseWriter
	status      int
	wroteHeader bool
	traceID     string
}

func (rec *statusRecorder) WriteHeader(code int) {
	if rec.wroteHeader == false {
		rec.wroteHeader = true
		rec.status = code
		rec.Header().Set(defaultTracingHTTPHeader, rec.traceID)
	}
	rec.ResponseWriter.WriteHeader(code)
}

func isIgnoredURL(url string, ignoredURLs []string) bool {
	for _, u := range ignoredURLs {
		if url == u {
			return true
		}
	}
	return false
}

func isIgnoredMethod(method string, ignoredMethods []string) bool {
	for _, m := range ignoredMethods {
		if method == m {
			return true
		}
	}
	return false
}

// TraceRequest wraps http.Handler and traces incoming request
func TraceRequest(tracer opentracing.Tracer, ignoredURLs []string,
	ignoredMethods []string) mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		handlerFunc := func(w http.ResponseWriter, r *http.Request) {
			url := r.URL.String()
			if isIgnoredURL(url, ignoredURLs) || isIgnoredMethod(r.Method, ignoredMethods) {
				next.ServeHTTP(w, r)
				return
			}

			ctx, _ := tracer.Extract(opentracing.HTTPHeaders, opentracing.HTTPHeadersCarrier(r.Header))
			functionName := r.Method + " " + r.URL.String()
			span := tracer.StartSpan(functionName, ext.RPCServerOption(ctx))
			ext.HTTPMethod.Set(span, r.Method)
			ext.HTTPUrl.Set(span, r.URL.String())
			ext.Component.Set(span, defaultComponentName)

			jaegerSpanContext := span.Context().(jaeger.SpanContext)
			traceID := jaegerSpanContext.String()

			responseWriter := &statusRecorder{w, http.StatusInternalServerError, false, traceID}
			r = r.WithContext(opentracing.ContextWithSpan(r.Context(), span))

			defer func() {
				ext.HTTPStatusCode.Set(span, uint16(responseWriter.status))
				if responseWriter.status >= http.StatusInternalServerError || !responseWriter.wroteHeader {
					ext.Error.Set(span, true)
				}
				span.Finish()
			}()

			next.ServeHTTP(responseWriter, r)
		}
		return http.HandlerFunc(handlerFunc)
	}
}
