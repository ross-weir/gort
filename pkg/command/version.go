package command

import "github.com/spf13/cobra"

func (r *Runner) initVersion() {
	cmd := &cobra.Command{
		Use:   "version",
		Short: "Output gort version",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Printf("gort has version %s built on %s\n", r.version, r.date)
		},
	}

	r.rootCmd.AddCommand(cmd)
}
