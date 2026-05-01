package validate

import (
	"strings"
	"testing"

	"alloy-config-generator/internal/config"
)

func TestInvalidLogsValue(t *testing.T) {
	cfg := config.DefaultPipelineConfig()
	cfg.Logs = "promtail"
	requireErrorContains(t, NormalizeAndValidate(&cfg), "logs")
}

func TestInvalidMetricsValue(t *testing.T) {
	cfg := config.DefaultPipelineConfig()
	cfg.Metrics = "victoriametrics"
	requireErrorContains(t, NormalizeAndValidate(&cfg), "metrics")
}

func TestInvalidTracesValue(t *testing.T) {
	cfg := config.DefaultPipelineConfig()
	cfg.Traces = "jaeger"
	requireErrorContains(t, NormalizeAndValidate(&cfg), "traces")
}

func TestInvalidProfile(t *testing.T) {
	cfg := config.DefaultPipelineConfig()
	cfg.Profile = "gold"
	requireErrorContains(t, NormalizeAndValidate(&cfg), "profile")
}

func TestInvalidMode(t *testing.T) {
	cfg := config.DefaultPipelineConfig()
	cfg.Mode = "operator"
	requireErrorContains(t, NormalizeAndValidate(&cfg), "mode")
}

func TestInvalidFormat(t *testing.T) {
	cfg := config.DefaultPipelineConfig()
	cfg.Format = "json"
	requireErrorContains(t, NormalizeAndValidate(&cfg), "format")
}

func TestAllPipelinesDisabledValidation(t *testing.T) {
	cfg := config.DefaultPipelineConfig()
	cfg.Logs = config.LogsNone
	cfg.Metrics = config.MetricsNone
	cfg.Traces = config.TracesNone
	requireErrorContains(t, NormalizeAndValidate(&cfg), "at least one pipeline")
}

func TestEndpointPlaceholdersAreAppliedWhenEmpty(t *testing.T) {
	cfg := config.DefaultPipelineConfig()
	cfg.LokiURL = ""
	cfg.RemoteWriteURL = ""
	cfg.TempoEndpoint = ""
	cfg.OTLPEndpoint = ""
	cfg.Traces = config.TracesOTLP

	if err := NormalizeAndValidate(&cfg); err != nil {
		t.Fatalf("NormalizeAndValidate returned error: %v", err)
	}
	if cfg.LokiURL != config.DefaultLokiURL {
		t.Fatalf("LokiURL = %q, want %q", cfg.LokiURL, config.DefaultLokiURL)
	}
	if cfg.RemoteWriteURL != config.DefaultRemoteWriteURL {
		t.Fatalf("RemoteWriteURL = %q, want %q", cfg.RemoteWriteURL, config.DefaultRemoteWriteURL)
	}
	if cfg.OTLPEndpoint != config.DefaultOTLPEndpoint {
		t.Fatalf("OTLPEndpoint = %q, want %q", cfg.OTLPEndpoint, config.DefaultOTLPEndpoint)
	}
}

func requireErrorContains(t *testing.T, err error, want string) {
	t.Helper()
	if err == nil {
		t.Fatalf("expected error containing %q, got nil", want)
	}
	if !strings.Contains(err.Error(), want) {
		t.Fatalf("error = %q, want to contain %q", err.Error(), want)
	}
}
