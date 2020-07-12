# Gort

![CI](https://github.com/ross-weir/gort/workflows/CI/badge.svg) [![Go Report Card](https://goreportcard.com/badge/github.com/ross-weir/gort)](https://goreportcard.com/report/github.com/ross-weir/gort) [![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

Gort is a simple port scanner used to detect and list open ports on a network host written in golang.

This project was my first time writing go. The goal was to learn about the language, tools, conventions and CI/CD processes of a go project.

## Improvements

Given the nature of the project some implementation is crude and lacking in features. Below is a non-exhaustive list of improvements that I would like to make in the future if time permits.

- Add more tests.
- Support multiple target hosts.
- Support UDP ports.
- Improved logging (verbose flag is unused, etc).
- Improved outputs/different output formats such as JSON (currently just prints to log).
- Better determine goroutine concurrency. We currently go by CPU count which probably isn't the best for this use-case.
- Fingerprint ports for service discovery.

## Usage

Don't run this against targets without permission. The author takes no responsibility for any damage caused.

For a list of available commands:

```
$ gort help

Simple port scanner. Run against an IP address to discover open ports.

Usage:
  gort [flags]
  gort [command]

Available Commands:
  help        Help about any command
  version     Output gort version

Flags:
      --all-ports               scan all ports
      --common-ports            scan common service ports
  -j, --concurrency int         concurrency (default NumCPU) (default 16)
  -h, --help                    help for gort
  -i, --ip string               target ip address
  -t, --port-timeout duration   how long to wait for a response from the port (default 1s)
  -p, --ports stringArray       comma separated list of target ports
  -v, --verbose                 use verbose output

Use "gort [command] --help" for more information about a command.
```

Scan for common ports open on our local machine:

`$ gort --ip 0.0.0.0 --common-ports`

Example output:

`Open port: 80`