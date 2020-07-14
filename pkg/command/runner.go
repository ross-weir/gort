package command

import (
	"github.com/ross-weir/gort/pkg/config"
	"github.com/spf13/cobra"
)

type Runner struct {
	rootCmd               *cobra.Command
	version, commit, date string
	cfg                   *config.Config
}

func NewRunner(version, commit, date string) *Runner {
	r := &Runner{
		version: version,
		commit:  commit,
		date:    date,
		cfg:     config.New(),
	}

	r.initRoot()
	r.initVersion()

	return r
}

func (r *Runner) Execute() error {
	return r.rootCmd.Execute()
}
