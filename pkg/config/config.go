package config

import "time"

type Config struct {
	IsVerbose   bool
	Concurrency int
	IP          string
	Ports       []string
	PortTimeout time.Duration
}

func NewDefault() *Config {
	return &Config{}
}
