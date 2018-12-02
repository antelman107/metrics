// Package cli implement util cli functions.
package cli

import (
	"fmt"
	"os"
)

// StopWithError print to stderr error message and stop application with
// non zero exit code
func StopWithError(err error) {
	fmt.Fprintf(os.Stderr, "Error: %v\n", err.Error())
	os.Exit(1)
}
