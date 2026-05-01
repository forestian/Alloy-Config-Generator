package generator

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"alloy-config-generator/internal/config"
)

func TestRenderConfigAlloyWithLogsPipeline(t *testing.T) {
	cfg := config.DefaultPipelineConfig()
	cfg.Metrics = config.MetricsNone
	cfg.Traces = config.TracesNone
	cfg.Format = config.FormatConfig

	content := renderNamedFile(t, &cfg, "config/config.alloy")

	requireContains(t, content, `loki.source.kubernetes "pods"`)
	requireContains(t, content, `loki.write "default"`)
	requireContains(t, content, config.DefaultLokiURL)
	requireContains(t, content, "Avoid high-cardinality labels")
}

func TestRenderConfigAlloyWithMetricsPipeline(t *testing.T) {
	cfg := config.DefaultPipelineConfig()
	cfg.Logs = config.LogsNone
	cfg.Traces = config.TracesNone
	cfg.Format = config.FormatConfig

	content := renderNamedFile(t, &cfg, "config/config.alloy")

	requireContains(t, content, `prometheus.scrape "kubernetes_pods"`)
	requireContains(t, content, `prometheus.remote_write "default"`)
	requireContains(t, content, config.DefaultRemoteWriteURL)
	requireContains(t, content, `cluster = "default-cluster"`)
}

func TestRenderConfigAlloyWithTracesPipeline(t *testing.T) {
	cfg := config.DefaultPipelineConfig()
	cfg.Logs = config.LogsNone
	cfg.Metrics = config.MetricsNone
	cfg.Format = config.FormatConfig

	content := renderNamedFile(t, &cfg, "config/config.alloy")

	requireContains(t, content, `otelcol.receiver.otlp "default"`)
	requireContains(t, content, `otelcol.processor.batch "default"`)
	requireContains(t, content, `otelcol.exporter.otlp "default"`)
	requireContains(t, content, config.DefaultTempoEndpoint)
}

func TestRenderHelmValues(t *testing.T) {
	cfg := config.DefaultPipelineConfig()
	cfg.Format = config.FormatHelm

	content := renderNamedFile(t, &cfg, "helm/values.yaml")

	requireContains(t, content, "configMap:")
	requireContains(t, content, "content: |-")
	requireContains(t, content, `discovery.kubernetes "pods"`)
	requireContains(t, content, "type: daemonset")
}

func TestOutputOverwriteProtection(t *testing.T) {
	cfg := config.DefaultPipelineConfig()
	cfg.OutputDir = t.TempDir()
	cfg.Format = config.FormatConfig

	if err := Generate(cfg); err != nil {
		t.Fatalf("first Generate returned error: %v", err)
	}
	if err := Generate(cfg); err == nil {
		t.Fatalf("second Generate returned nil error, want overwrite protection")
	} else if !strings.Contains(err.Error(), "refusing to overwrite") {
		t.Fatalf("error = %q, want overwrite protection", err.Error())
	}

	cfg.Force = true
	if err := Generate(cfg); err != nil {
		t.Fatalf("Generate with Force returned error: %v", err)
	}
}

func TestInitOutputStructure(t *testing.T) {
	cfg := config.DefaultPipelineConfig()
	cfg.OutputDir = t.TempDir()
	cfg.Format = config.FormatAll

	if err := Generate(cfg); err != nil {
		t.Fatalf("Generate returned error: %v", err)
	}

	for _, path := range []string{
		"README.md",
		"config/config.alloy",
		"helm/values.yaml",
		"examples/install.sh",
		"examples/uninstall.sh",
	} {
		fullPath := filepath.Join(cfg.OutputDir, filepath.FromSlash(path))
		if _, err := os.Stat(fullPath); err != nil {
			t.Fatalf("expected generated file %s: %v", fullPath, err)
		}
	}
}

func renderNamedFile(t *testing.T, cfg *config.PipelineConfig, name string) string {
	t.Helper()
	files, err := RenderFiles(cfg)
	if err != nil {
		t.Fatalf("RenderFiles returned error: %v", err)
	}
	for _, file := range files {
		if file.Path == name {
			return file.Content
		}
	}
	t.Fatalf("generated file %q not found", name)
	return ""
}

func requireContains(t *testing.T, content, want string) {
	t.Helper()
	if !strings.Contains(content, want) {
		t.Fatalf("content does not contain %q:\n%s", want, content)
	}
}
