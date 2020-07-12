package command

import (
	"runtime"
	"time"

	"github.com/ross-weir/gort/pkg/config"
	"github.com/ross-weir/gort/pkg/portscan"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

func (r *Runner) initRoot() {
	cmd := &cobra.Command{
		Use:   "gort",
		Short: "gort is a simple port scanner.",
		Long:  "Simple port scanner. Run against an IP address to discover open ports.",
		Run: func(cmd *cobra.Command, args []string) {
			portscan.ScanPorts(r.cfg)
		},
	}

	initRootFlagSet(cmd.Flags(), r.cfg)
	markRequiredFlags(cmd)
	r.rootCmd = cmd
}

func initRootFlagSet(fs *pflag.FlagSet, cfg *config.Config) {
	fs.BoolVarP(&cfg.IsVerbose, "verbose", "v", false, "use verbose output")
	fs.BoolVar(&cfg.AllPorts, "all-ports", false, "scan all ports")
	fs.BoolVar(&cfg.CommonPorts, "common-ports", false, "scan common service ports")
	fs.IntVarP(&cfg.Concurrency, "concurrency", "j", runtime.NumCPU(), "concurrency (default NumCPU)")
	fs.StringVarP(&cfg.IP, "ip", "i", "", "target ip address")
	fs.StringSliceVarP(&cfg.Ports, "ports", "p", []string{}, "comma separated list of target ports")
	fs.DurationVarP(&cfg.PortTimeout, "port-timeout", "t", time.Second, "how long to wait for a response from the port")
}

func markRequiredFlags(cmd *cobra.Command) {
	_ = cmd.MarkFlagRequired("ip")
}
