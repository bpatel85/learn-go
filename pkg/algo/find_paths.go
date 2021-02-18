package algo

type FindPaths struct{}

func (fp FindPaths) numPathsInner(world [][]int, x int, y int) int {
	if world[x][y] == 1 {
		return 0
	}

	if x == 0 || y == 0 {
		return 1
	}

	return fp.numPathsInner(world, x-1, y) + fp.numPathsInner(world, x, y-1)
}

func (fp FindPaths) CountNumPaths(world [][]int) int {
	// add some checks
	if world == nil {
		return 0
	}

	return fp.numPathsInner(world, len(world)-1, len(world[0])-1)
}
