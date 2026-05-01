package cmd

import "github.com/spf13/cobra"

var Version = "0.1.0"

func NewRootCommand(version string) *cobra.Command {
	root := &cobra.Command{
		Use:           "alloygen",
		Short:         "Generate starter Grafana Alloy configuration",
		SilenceUsage:  true,
		SilenceErrors: true,
	}

	root.AddCommand(newInitCommand())
	root.AddCommand(newGenerateCommand())
	root.AddCommand(newVersionCommand(version))

	return root
}

func Execute() error {
	return NewRootCommand(Version).Execute()
}
