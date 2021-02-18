package algo_test

import (
	"testing"
	"github.com/bpatel85/learn-go/pkg/algo"
)
type FindPathTestStructs struct {
	input [][]int
	expected int
}

func TestNumPaths(t *testing.T) {
	testRuns := []FindPathTestStructs {
		{
			input: [][]int{
				{0, 0, 0},
				{0, 0, 0},
				{0, 0, 0},
			},
			expected: 6,
		},
		{
			input: [][]int{
				{0},
				{0},
				{0},
			},
			expected: 1,
		},
		{
			input: [][]int{
				{0},
			},
			expected: 1,
		},
		{
			input: nil,
			expected: 0,
		},
		{
			input: [][]int{
				{0, 0, 0},
				{0, 1, 0},
				{0, 0, 0},
			},
			expected: 2,
		},
	}
	
	fp := algo.FindPaths{}

	for _, tc := range testRuns {
		actual := fp.CountNumPaths(tc.input)
		if actual != tc.expected {
			t.Errorf("not right number of path. Expected: %d, Actual: %d", tc.expected, actual )
		}
	}
}
