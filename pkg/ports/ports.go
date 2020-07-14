package ports

import (
	"fmt"
	"net"
	"sync"
	"time"

	"github.com/ross-weir/gort/pkg/config"
)

type scanRoutineArg struct {
	ip      string
	ports   []string
	timeout time.Duration
	c       chan string
	wg      *sync.WaitGroup
}

// TODO: If concurrency is equal to or higher than len(ports) then return slice of slices each containing one port.
// Currently all ports are assigned to a single goroutine, it's less efficient.
func dividePorts(ports []string, concurrency int) [][]string {
	var divided [][]string

	for i := 0; i < len(ports); i += concurrency {
		end := i + concurrency

		if end > len(ports) {
			end = len(ports)
		}

		divided = append(divided, ports[i:end])
	}

	return divided
}

func checkPorts(arg *scanRoutineArg) {
	for _, port := range arg.ports {
		conn, _ := net.DialTimeout("tcp", net.JoinHostPort(arg.ip, port), arg.timeout)
		if conn != nil {
			conn.Close()
			arg.c <- port
		}
	}

	arg.wg.Done()
}

func Scan(cfg *config.Config) {
	dividedPorts := dividePorts(cfg.Ports(), cfg.Concurrency)
	c := make(chan string)
	wg := &sync.WaitGroup{}

	for _, ports := range dividedPorts {
		wg.Add(1)
		go checkPorts(&scanRoutineArg{
			cfg.IP,
			ports,
			cfg.PortTimeout,
			c,
			wg,
		})
	}

	go func() {
		wg.Wait()
		close(c)
	}()

	for msg := range c {
		fmt.Printf("Open port: %s\n", msg)
	}
}
