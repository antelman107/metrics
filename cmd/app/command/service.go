// Package command contains cli commands.
package command

import (
	"errors"
	"os"
	"os/signal"
	"runtime"
	"syscall"

	"github.com/antelman107/metrics/container"
	"github.com/spf13/cobra"

	"github.com/antelman107/metrics/cmd/app/definition/service"
)

// Service command.
var serviceCmd = &cobra.Command{
	Use:   "service <name>",
	Short: "Run various services",
	Args:  cobra.ExactArgs(1),
	RunE:  serviceCmdListener,
}

// Command init function.
func init() {
	rootCmd.AddCommand(serviceCmd)

	serviceCmd.Flags().IntP("goroutines", "g", runtime.NumCPU(), "Goroutines num")
	serviceCmd.Flags().StringP("worker", "w", "1", "Worker name")
}

// Command handler func.
func serviceCmdListener(cmd *cobra.Command, args []string) (err error) {
	// validation
	if err = cobra.ExactArgs(1)(cmd, args); err != nil {
		return err
	}

	cmd.ValidArgs = make([]string, 0, 8)

	var serviceDefs = make(map[string]string, 8)
	if err = container.Iterate(diContainer, service.DefServiceTag, func(ctx container.Context, tag *container.Tag, name string) error {
		if _, exist := tag.Args["name"]; !exist {
			return errors.New("definition should have name argument")
		}

		serviceDefs[tag.Args["name"]] = name
		cmd.ValidArgs = append(cmd.ValidArgs, tag.Args["name"])

		return nil
	}); err != nil {
		return err
	}

	if len(cmd.ValidArgs) == 0 {
		return errors.New("can't find any service")
	}

	if err = cobra.OnlyValidArgs(cmd, args); err != nil {
		return err
	}

	// dispatch
	var instance service.Service
	if err = diContainer.Fill(serviceDefs[args[0]], &instance); err != nil {
		return err
	}

	var goroutines int
	if goroutines, err = cmd.LocalFlags().GetInt("goroutines"); err != nil {
		return err
	}

	var signalChan = make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)

	if err = instance.Start(goroutines); err != nil {
		return err
	}

	select {
	case <-signalChan:
	case <-instance.Done():
		return nil
	}

	return instance.Stop()
}
