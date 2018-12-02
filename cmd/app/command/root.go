// Package command contains cli commands.
package command

import (
	"os"
	"path/filepath"

	"github.com/spf13/cobra"

	"github.com/antelman107/metrics/cli/command"
	"github.com/antelman107/metrics/container"
)

var (
	// Config path.
	configPath string

	// DI Container.
	diContainer container.Context

	// Root command.
	rootCmd = &cobra.Command{
		Use:           "app [command]",
		SilenceUsage:  true,
		SilenceErrors: true,
		PersistentPreRunE: func(cmd *cobra.Command, args []string) (err error) {

			diContainer, err = container.Instance(map[string]interface{}{
				"cli_cmd":     cmd,
				"cli_args":    args,
				"config_path": configPath,
			})

			return err
		},
	}
)

// Execute root cmd.
func Execute() (err error) {
	var appPath string
	if appPath, err = filepath.Abs(filepath.Dir(os.Args[0])); err != nil {
		return err
	}

	// Application config path
	rootCmd.PersistentFlags().StringVarP(&configPath, "config", "c", appPath+"/config.json", "arConfig file")

	// Append help command
	command.RegisterCmd(rootCmd)

	// Run
	err = rootCmd.Execute()

	// Delete context
	if diContainer != nil {
		diContainer.Delete()
	}

	return err
}
