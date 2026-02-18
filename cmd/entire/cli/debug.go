package cli

import "github.com/spf13/cobra"

func newDebugCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:    "debug",
		Short:  "Debug commands for troubleshooting",
		Hidden: true, // Hidden from help output
		RunE: func(cmd *cobra.Command, _ []string) error {
			return cmd.Help()
		},
	}

	return cmd
}
