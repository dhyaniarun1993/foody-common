package tracer

import (
	"fmt"
	"io"

	"github.com/opentracing/opentracing-go"

	"github.com/uber/jaeger-client-go"
	config "github.com/uber/jaeger-client-go/config"
)

// InitJaeger creates and set global Tracer
func InitJaeger(configuration Configuration) (opentracing.Tracer, io.Closer) {
	cfg := &config.Configuration{
		ServiceName: configuration.ServiceName,
		RPCMetrics:  configuration.RPCMetrics,
		Sampler: &config.SamplerConfig{
			Type:  configuration.Sampler.Type,
			Param: configuration.Sampler.Param,
		},
		Reporter: &config.ReporterConfig{
			LogSpans:          configuration.Reporter.LogSpans,
			CollectorEndpoint: configuration.Reporter.CollectorEndpoint,
		},
	}
	tracer, closer, err := cfg.NewTracer(config.Logger(jaeger.StdLogger))
	if err != nil {
		panic(fmt.Sprintf("ERROR: cannot init Jaeger: %v\n", err))
	}
	return tracer, closer
}
