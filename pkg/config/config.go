package config

import (
	"time"
)

type Config struct {
	IsVerbose   bool
	Concurrency int
	IP          string
	Ports       []string
	PortTimeout time.Duration
	AllPorts    bool
	CommonPorts bool
}

func NewDefault() *Config {
	return &Config{}
}
