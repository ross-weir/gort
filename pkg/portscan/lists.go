package portscan

import "strconv"

const maxPort = 65536

func calculateAllPorts() []string {
	var results []string

	for i := 1; i <= maxPort; i++ {
		results = append(results, strconv.Itoa(i))
	}

	return results
}

var AllPorts = calculateAllPorts()
