package config

import "testing"

func TestDefaultPipelineConfig(t *testing.T) {
	cfg := DefaultPipelineConfig()

	if cfg.OutputDir != DefaultOutputDir {
		t.Fatalf("OutputDir = %q, want %q", cfg.OutputDir, DefaultOutputDir)
	}
	if cfg.Namespace != DefaultNamespace {
		t.Fatalf("Namespace = %q, want %q", cfg.Namespace, DefaultNamespace)
	}
	if cfg.ClusterName != DefaultClusterName {
		t.Fatalf("ClusterName = %q, want %q", cfg.ClusterName, DefaultClusterName)
	}
	if cfg.Logs != LogsLoki || cfg.Metrics != MetricsMimir || cfg.Traces != TracesTempo {
		t.Fatalf("pipelines = %q/%q/%q, want loki/mimir/tempo", cfg.Logs, cfg.Metrics, cfg.Traces)
	}
	if cfg.Profile != ProfileDev || cfg.Mode != ModeKubernetes || cfg.Format != FormatAll {
		t.Fatalf("profile/mode/format = %q/%q/%q, want dev/kubernetes/all", cfg.Profile, cfg.Mode, cfg.Format)
	}
	if cfg.LokiURL != DefaultLokiURL || cfg.RemoteWriteURL != DefaultRemoteWriteURL || cfg.TempoEndpoint != DefaultTempoEndpoint || cfg.OTLPEndpoint != DefaultOTLPEndpoint {
		t.Fatalf("endpoint defaults were not populated")
	}
}
