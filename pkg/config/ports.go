package config

import "strconv"

const maxPort = 65536

func calculateAllPorts() []string {
	var results []string

	for i := 1; i <= maxPort; i++ {
		results = append(results, strconv.Itoa(i))
	}

	return results
}

var allPorts = calculateAllPorts()

var commonPorts = []string{
	"20", "21", "22", "23", "25", "53", "67", "68", "69", "80", "110", "119", "123",
	"135", "136", "137", "138", "139", "143", "161", "162", "389", "443", "3389",
}
