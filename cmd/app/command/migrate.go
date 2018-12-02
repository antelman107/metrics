// Package command contains cli commands.
package command

import (
	"github.com/spf13/cobra"
)

// Migrate command.
var migrateCmd = &cobra.Command{
	Use:   "migrate [command]",
	Short: "Migrate commands group",
	Args:  cobra.ExactArgs(1),
}

// Command init function.
func init() {
	rootCmd.AddCommand(migrateCmd)
}
