// Package main contains main fuc of application.
package main

import (
	"os"

	"github.com/antelman107/metrics/cli"
	"github.com/antelman107/metrics/cmd/app/command"
)

func main() {
	if err := command.Execute(); err != nil {
		cli.StopWithError(err)
	}

	os.Exit(0)
}
