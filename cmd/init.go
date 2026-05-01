package cmd

import (
	"fmt"

	"alloy-config-generator/internal/config"
	"alloy-config-generator/internal/generator"

	"github.com/spf13/cobra"
)

func newInitCommand() *cobra.Command {
	cfg := config.DefaultPipelineConfig()
	cfg.Format = config.FormatAll

	cmd := &cobra.Command{
		Use:   "init",
		Short: "Create an example Alloy generator project directory",
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := generator.Generate(cfg); err != nil {
				return err
			}

			fmt.Fprintf(cmd.OutOrStdout(), "Created Alloy starter project in %s\n", cfg.OutputDir)
			return nil
		},
	}

	cmd.Flags().StringVar(&cfg.OutputDir, "output", cfg.OutputDir, "output directory")
	cmd.Flags().BoolVar(&cfg.Force, "force", cfg.Force, "overwrite generated files if they already exist")

	return cmd
}
