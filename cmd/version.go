package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func newVersionCommand(version string) *cobra.Command {
	return &cobra.Command{
		Use:   "version",
		Short: "Print alloygen version",
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Fprintf(cmd.OutOrStdout(), "alloygen version %s\n", version)
			return nil
		},
	}
}
