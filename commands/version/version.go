package version

import (
	"fmt"

	"github.com/spf13/cobra"
)

const (
	Version = "0.1.0-pre"
)

func VersionCommandHandler(cmd *cobra.Command, args []string) {
	fmt.Printf("rbtxt\nversion : %s\n", Version)
}
