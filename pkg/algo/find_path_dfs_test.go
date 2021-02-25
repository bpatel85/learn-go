package algo_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/bpatel85/learn-go/pkg/algo"
)

/**
Input:
[
  [0,0,1],
  [0,0,0],
  [0,0,0]
]

Expected Output: 7
Paths:
Down -> Down  -> Right -> Right
Down -> Right -> Right -> Down
Down -> Right -> Down -> Right
Down -> Down -> Right -> Up -> Right -> Down
Right -> Down -> Right -> Down
Right -> Down -> Down -> Right
Right -> Down -> Left -> Down -> Right -> Right
**/

func TestFindAllPaths_DFS(t *testing.T) {
	maze := [][]int{
		{0, 0, 1},
		{0, 0, 0},
		{0, 0, 0},
	}

	paths, err := algo.FindPath_DFS(maze, 0, 0, 2, 2)
	if err != nil {
		t.Errorf(err.Error())
	}

	fmt.Printf("\nthere are %d paths.\n", len(paths))
	for _, path := range paths {
		fmt.Printf("\n%v", strings.Join(path, " -> "))
	}

	if 7 != len(paths) {
		t.Errorf("error in expected number of paths")
	}
}
