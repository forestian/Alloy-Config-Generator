package cmd

import (
	"fmt"

	"alloy-config-generator/internal/config"
	"alloy-config-generator/internal/generator"

	"github.com/spf13/cobra"
)

func newGenerateCommand() *cobra.Command {
	cfg := config.DefaultPipelineConfig()

	cmd := &cobra.Command{
		Use:   "generate",
		Short: "Generate Grafana Alloy config and optional Helm values",
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := generator.Generate(cfg); err != nil {
				return err
			}

			fmt.Fprintf(cmd.OutOrStdout(), "Generated Alloy starter files in %s\n", cfg.OutputDir)
			return nil
		},
	}

	flags := cmd.Flags()
	flags.StringVar(&cfg.OutputDir, "output", cfg.OutputDir, "output directory")
	flags.StringVar(&cfg.Namespace, "namespace", cfg.Namespace, "Kubernetes namespace for Alloy")
	flags.StringVar(&cfg.ClusterName, "cluster-name", cfg.ClusterName, "cluster label attached to generated pipelines")
	flags.StringVar(&cfg.Logs, "logs", cfg.Logs, "logs backend: none or loki")
	flags.StringVar(&cfg.Metrics, "metrics", cfg.Metrics, "metrics backend: none, mimir, or prometheus")
	flags.StringVar(&cfg.Traces, "traces", cfg.Traces, "traces backend: none, tempo, or otlp")
	flags.StringVar(&cfg.RemoteWriteURL, "remote-write-url", cfg.RemoteWriteURL, "Prometheus remote_write endpoint")
	flags.StringVar(&cfg.LokiURL, "loki-url", cfg.LokiURL, "Loki push endpoint")
	flags.StringVar(&cfg.TempoEndpoint, "tempo-endpoint", cfg.TempoEndpoint, "Tempo OTLP gRPC endpoint")
	flags.StringVar(&cfg.OTLPEndpoint, "otlp-endpoint", cfg.OTLPEndpoint, "generic OTLP gRPC endpoint")
	flags.StringVar(&cfg.Profile, "profile", cfg.Profile, "profile: dev or production")
	flags.StringVar(&cfg.Mode, "mode", cfg.Mode, "mode: kubernetes or standalone")
	flags.StringVar(&cfg.Format, "format", cfg.Format, "format: config, helm, or all")
	flags.BoolVar(&cfg.Force, "force", cfg.Force, "overwrite generated files if they already exist")

	return cmd
}
