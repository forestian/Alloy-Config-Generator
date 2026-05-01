package validate

import (
	"fmt"
	"strings"

	"alloy-config-generator/internal/config"
)

func NormalizeAndValidate(cfg *config.PipelineConfig) error {
	trimFields(cfg)

	if cfg.OutputDir == "" {
		return fmt.Errorf("output must not be empty")
	}
	if cfg.Namespace == "" {
		return fmt.Errorf("namespace must not be empty")
	}
	if cfg.ClusterName == "" {
		return fmt.Errorf("cluster-name must not be empty")
	}
	if !allowed(cfg.Logs, config.LogsNone, config.LogsLoki) {
		return fmt.Errorf("logs must be one of: none, loki")
	}
	if !allowed(cfg.Metrics, config.MetricsNone, config.MetricsMimir, config.MetricsPrometheus) {
		return fmt.Errorf("metrics must be one of: none, mimir, prometheus")
	}
	if !allowed(cfg.Traces, config.TracesNone, config.TracesTempo, config.TracesOTLP) {
		return fmt.Errorf("traces must be one of: none, tempo, otlp")
	}
	if !allowed(cfg.Profile, config.ProfileDev, config.ProfileProduction) {
		return fmt.Errorf("profile must be one of: dev, production")
	}
	if !allowed(cfg.Mode, config.ModeKubernetes, config.ModeStandalone) {
		return fmt.Errorf("mode must be one of: kubernetes, standalone")
	}
	if !allowed(cfg.Format, config.FormatConfig, config.FormatHelm, config.FormatAll) {
		return fmt.Errorf("format must be one of: config, helm, all")
	}
	if cfg.Logs == config.LogsNone && cfg.Metrics == config.MetricsNone && cfg.Traces == config.TracesNone {
		return fmt.Errorf("at least one pipeline must be enabled")
	}

	if cfg.Logs == config.LogsLoki && cfg.LokiURL == "" {
		cfg.LokiURL = config.DefaultLokiURL
	}
	if cfg.HasMetrics() && cfg.RemoteWriteURL == "" {
		cfg.RemoteWriteURL = config.DefaultRemoteWriteURL
	}
	if cfg.Traces == config.TracesTempo && cfg.TempoEndpoint == "" {
		cfg.TempoEndpoint = config.DefaultTempoEndpoint
	}
	if cfg.Traces == config.TracesOTLP && cfg.OTLPEndpoint == "" {
		cfg.OTLPEndpoint = config.DefaultOTLPEndpoint
	}

	return nil
}

func trimFields(cfg *config.PipelineConfig) {
	cfg.OutputDir = strings.TrimSpace(cfg.OutputDir)
	cfg.Namespace = strings.TrimSpace(cfg.Namespace)
	cfg.ClusterName = strings.TrimSpace(cfg.ClusterName)
	cfg.Logs = strings.TrimSpace(strings.ToLower(cfg.Logs))
	cfg.Metrics = strings.TrimSpace(strings.ToLower(cfg.Metrics))
	cfg.Traces = strings.TrimSpace(strings.ToLower(cfg.Traces))
	cfg.LokiURL = strings.TrimSpace(cfg.LokiURL)
	cfg.RemoteWriteURL = strings.TrimSpace(cfg.RemoteWriteURL)
	cfg.TempoEndpoint = strings.TrimSpace(cfg.TempoEndpoint)
	cfg.OTLPEndpoint = strings.TrimSpace(cfg.OTLPEndpoint)
	cfg.Profile = strings.TrimSpace(strings.ToLower(cfg.Profile))
	cfg.Mode = strings.TrimSpace(strings.ToLower(cfg.Mode))
	cfg.Format = strings.TrimSpace(strings.ToLower(cfg.Format))
}

func allowed(value string, values ...string) bool {
	for _, allowedValue := range values {
		if value == allowedValue {
			return true
		}
	}
	return false
}
