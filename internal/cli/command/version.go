package command

import (
	"fmt"

	"github.com/aqyuki/rbtxt/internal/info"
	"github.com/spf13/cobra"
)

func CreateVersionCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "version",
		Short: "Print the version number of rbtxt",
		Long:  `All software has versions. This is rbtxt's`,
		Run: func(cmd *cobra.Command, args []string) {
			info, err := info.GetApplicationInformation()
			if err != nil {
				fmt.Printf("failed to load application information. Please report develop team.")
				return
			}
			fmt.Printf("rbtxt\nversion : %s\n", info.Version)
		},
	}
}
