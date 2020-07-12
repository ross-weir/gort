package main

import (
	"fmt"
	"os"

	"github.com/ross-weir/gort/pkg/command"
	"github.com/ross-weir/gort/pkg/exitcode"
)

var (
	// Populated by goreleaser during build
	version = "master"
	commit  = "none"
	date    = "unknown"
)

func main() {
	r := command.NewRunner(version, commit, date)

	if err := r.Execute(); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "failed executing command with error %v\n", err)
		os.Exit(exitcode.Failure)
	}
}
