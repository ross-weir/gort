package portscan

import (
	"testing"

	assert "github.com/stretchr/testify/require"
)

func TestDividePorts(t *testing.T) {
	testCases := []struct {
		allPorts    []string
		concurrency int
		expected    [][]string
	}{
		{
			[]string{"21", "80", "443"},
			3,
			[][]string{
				{"21", "80", "443"},
			},
		},
		// Small odd list.
		{
			[]string{"21", "80", "443"},
			2,
			[][]string{
				{"21", "80"},
				{"443"},
			},
		},
		// Small even list.
		{
			[]string{"21", "22", "80", "443"},
			2,
			[][]string{
				{"21", "22"},
				{"80", "443"},
			},
		},
		{
			[]string{"21", "22", "80", "443", "8080", "8081", "554", "271", "81", "661"},
			3,
			[][]string{
				{"21", "22", "80"},
				{"443", "8080", "8081"},
				{"554", "271", "81"},
				{"661"},
			},
		},
	}

	for _, tc := range testCases {
		actual := dividePorts(tc.allPorts, tc.concurrency)

		assert.ElementsMatch(t, actual, tc.expected)
	}
}