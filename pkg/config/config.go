package config

import (
	"time"
)

type Config struct {
	IsVerbose     bool
	Concurrency   int
	IP            string
	SuppliedPorts []string
	PortTimeout   time.Duration
	AllPorts      bool
	CommonPorts   bool
}

func (c *Config) Ports() []string {
	if c.AllPorts {
		return allPorts
	}

	if c.CommonPorts {
		return commonPorts
	}

	return c.SuppliedPorts
}

func New() *Config {
	return &Config{}
}
