package tracer

// SamplerConfig allows initializing a non-default sampler
type SamplerConfig struct {
	Type  string  `required:"true" split_words:"true"`
	Param float64 `required:"true" split_words:"true"`
}

// ReporterConfig configures the reporter.
type ReporterConfig struct {
	LogSpans          bool   `required:"true" split_words:"true"`
	CollectorEndpoint string `required:"true" split_words:"true"`
}

// Configuration provides config option to create Tracer
type Configuration struct {
	ServiceName string `required:"true" split_words:"true"`
	RPCMetrics  bool   `require:"true" split_words:"true"`
	Sampler     SamplerConfig
	Reporter    ReporterConfig
}
