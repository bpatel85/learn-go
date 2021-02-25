package algo

/**
Given a maze 2D MxN array with 1s and 0s, find number of paths from start of the maze (left upper corner) to end of the maze (right bottom corner).
Rules of travesal:
- you can only traverse right or down.
- if the cell value is set to 1, you cannot traverse through it.
**/
type FindNumPaths struct{}

// this is dynamic programming with recursion
func (fp FindNumPaths) numPathsInner(world [][]int, x int, y int) int {
	if world[x][y] == 1 {
		return 0
	}

	if x == 0 || y == 0 {
		return 1
	}

	return fp.numPathsInner(world, x-1, y) + fp.numPathsInner(world, x, y-1)
}

func (fp FindNumPaths) CountNumPaths(world [][]int) int {
	// add some checks
	if world == nil {
		return 0
	}

	return fp.numPathsInner(world, len(world)-1, len(world[0])-1)
}
