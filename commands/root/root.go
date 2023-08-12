package root

import "github.com/spf13/cobra"

func RootCommandHandler(cmd *cobra.Command, args []string) {
	cmd.Help()
	return
}
