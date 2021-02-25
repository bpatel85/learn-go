/**
MAZE = [[0,0,0,0,0],
        [1,0,1,0,1],
        [0,0,0,0,0],
        [0,1,1,1,1],
        [0,0,0,0,9]]
The mazes are passed in as a two dimensional array. 0's represent spaces you may traverse. 1's represent walls. The program always starts in the 0,0 coordinate of the maze which is guaranteed to be a 0. The 9 represents the end of the maze and may be positioned anywhere. In this example maze it lives at 4,4. You may only travel up, down, left and right, no diagonal moves are allowed. The maze has already been passed through a maze validator so it is guaranteed not to contain any bad types or bad values. No letter 'a' or -1 will make it to your program.

The success condition is to return a list of coordinates that would let us retrace a valid path through the maze.
There may not be a valid path for every maze, in which case you should return a false value like an empty set or a None or a False.

If you complete this program we might start to talk about the shortest path through the maze or the shortest running time solution, but to start any path through the maze is good. We might also later talk about very very large mazes, not just 5x5 mazes, but milliion x million element mazes.
**/

package algo

import (
	"fmt"

	"github.com/bpatel85/learn-go/pkg/common"
)

type traverseNode struct {
	x, y      int
	pathSoFar [][]int
	visitMap  [][]bool
	hops      []string
}

func isVisited(currentX int, currentY int, path [][]int) bool {
	for _, pt := range path {
		if pt[0] == currentX && pt[1] == currentY {
			return true
		}
	}

	return false
}

func FindPath_DFS(maze [][]int, startX, startY, endX, endY int) ([][]string, error) {
	// validate that start and end have valid vals
	// assume: maze validations is done
	visitMap := make([][]bool, len(maze))
	for i := range maze {
		visitMap[i] = make([]bool, len(maze[i]))
	}

	hops := []string{"left", "right", "up", "down"}
	allPaths := make([][]string, 0)
	stack := common.NewThreadSaftStack()
	startNode := traverseNode{
		x:         startX,
		y:         startY,
		pathSoFar: [][]int{{startX, startY}},
		visitMap:  visitMap,
	}
	stack.Push(startNode)

	for !stack.IsEmpty() {
		val, err := stack.Pop()
		if err != nil {
			return nil, err
		}
		node := val.(traverseNode)

		// check end condition
		if node.x == endX && node.y == endY {
			allPaths = append(allPaths, node.hops)
		}

		neighbours := [][]int{
			{node.x - 1, node.y}, // left
			{node.x + 1, node.y}, // right
			{node.x, node.y - 1}, // up
			{node.x, node.y + 1}, // down
		}

		for idx, n := range neighbours {
			// boundary condition
			if n[0] < 0 || n[0] >= len(maze[0]) || n[1] < 0 || n[1] >= len(maze[0]) {
				continue
			}

			if maze[n[0]][n[1]] == 1 || isVisited(n[0], n[1], node.pathSoFar) {
				continue
			}

			fmt.Printf("\nGoing to traves %s to (%d, %d) from (%d, %d)\n", hops[idx], n[0], n[1], node.x, node.y)
			stack.Push(traverseNode{
				x:         n[0],
				y:         n[1],
				pathSoFar: append(node.pathSoFar, []int{n[0], n[1]}),
				hops:      append(node.hops, hops[idx]),
			})

		}
	}

	return allPaths, nil
}
