package version

import (
	"fmt"

	"github.com/spf13/cobra"
)

const (
	Version = "0.1.1"
)

func VersionCommandHandler(cmd *cobra.Command, args []string) {
	fmt.Printf("rbtxt\nversion : %s\n", Version)
}
