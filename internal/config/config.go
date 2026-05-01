package config

const (
	LogsNone = "none"
	LogsLoki = "loki"

	MetricsNone       = "none"
	MetricsMimir      = "mimir"
	MetricsPrometheus = "prometheus"

	TracesNone  = "none"
	TracesTempo = "tempo"
	TracesOTLP  = "otlp"

	ProfileDev        = "dev"
	ProfileProduction = "production"

	ModeKubernetes = "kubernetes"
	ModeStandalone = "standalone"

	FormatConfig = "config"
	FormatHelm   = "helm"
	FormatAll    = "all"

	DefaultOutputDir      = "./alloy-generated"
	DefaultNamespace      = "monitoring"
	DefaultClusterName    = "default-cluster"
	DefaultLokiURL        = "http://loki-gateway.monitoring.svc:80/loki/api/v1/push"
	DefaultRemoteWriteURL = "http://mimir-nginx.monitoring.svc:80/api/v1/push"
	DefaultTempoEndpoint  = "tempo-distributor.monitoring.svc:4317"
	DefaultOTLPEndpoint   = "otel-collector.monitoring.svc:4317"
)

type PipelineConfig struct {
	OutputDir      string
	Namespace      string
	ClusterName    string
	Logs           string
	Metrics        string
	Traces         string
	LokiURL        string
	RemoteWriteURL string
	TempoEndpoint  string
	OTLPEndpoint   string
	Profile        string
	Mode           string
	Format         string
	Force          bool
}

func DefaultPipelineConfig() PipelineConfig {
	return PipelineConfig{
		OutputDir:      DefaultOutputDir,
		Namespace:      DefaultNamespace,
		ClusterName:    DefaultClusterName,
		Logs:           LogsLoki,
		Metrics:        MetricsMimir,
		Traces:         TracesTempo,
		LokiURL:        DefaultLokiURL,
		RemoteWriteURL: DefaultRemoteWriteURL,
		TempoEndpoint:  DefaultTempoEndpoint,
		OTLPEndpoint:   DefaultOTLPEndpoint,
		Profile:        ProfileDev,
		Mode:           ModeKubernetes,
		Format:         FormatAll,
	}
}

func (c PipelineConfig) HasLogs() bool {
	return c.Logs == LogsLoki
}

func (c PipelineConfig) HasMetrics() bool {
	return c.Metrics == MetricsMimir || c.Metrics == MetricsPrometheus
}

func (c PipelineConfig) HasTraces() bool {
	return c.Traces == TracesTempo || c.Traces == TracesOTLP
}

func (c PipelineConfig) IsProduction() bool {
	return c.Profile == ProfileProduction
}

func (c PipelineConfig) IsKubernetes() bool {
	return c.Mode == ModeKubernetes
}

func (c PipelineConfig) IncludesConfig() bool {
	return c.Format == FormatConfig || c.Format == FormatAll
}

func (c PipelineConfig) IncludesHelm() bool {
	return c.Format == FormatHelm || c.Format == FormatAll
}

func (c PipelineConfig) IncludesExamples() bool {
	return c.Format == FormatAll
}

func (c PipelineConfig) TraceEndpoint() string {
	if c.Traces == TracesOTLP {
		return c.OTLPEndpoint
	}
	return c.TempoEndpoint
}

func (c PipelineConfig) TraceBackendName() string {
	if c.Traces == TracesOTLP {
		return "OTLP-compatible backend"
	}
	return "Tempo"
}
