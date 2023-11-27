package command

import (
	"fmt"

	"github.com/spf13/cobra"
)

const (
	Version = "0.1.2"
)

func CreateVersionCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "version",
		Short: "Print the version number of rbtxt",
		Long:  `All software has versions. This is rbtxt's`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("rbtxt\nversion : %s\n", Version)
		},
	}
}
