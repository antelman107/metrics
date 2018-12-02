// Package command contains common cli commands.
package command

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	// Version
	appVersion = "unknown"

	// Version command
	versionCmd = &cobra.Command{
		Use:           "version",
		Short:         "Application version",
		SilenceUsage:  true,
		SilenceErrors: true,
		Run:           versionCmdHandler,
	}
)

// RegisterCmd add sub command version
func RegisterCmd(root *cobra.Command) {
	root.AddCommand(versionCmd)
}

// Version command handler
func versionCmdHandler(_ *cobra.Command, _ []string) {
	fmt.Println(appVersion)
}
