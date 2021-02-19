package algo

/**
Connected Components

input:
[ 1 1 1 1 1 1 ]
[ 1 0 0 0 0 1 ]
[ 1 0 1 1 0 1 ]
[ 1 0 1 1 0 1 ]
[ 1 0 0 0 0 1 ]
[ 1 1 1 1 1 1 ]

starting index: for example, [2, 2]

output:
lists of indices representing the components connected to that location in the grid
for this grid: [[2, 2], [2, 3], [3, 2], [3, 3]]
**/

type Cell struct {
	X, Y int
}

func FindConnectedComponents(world [][]int, start Cell) []Cell {
	connected := make([]Cell, 0)

	return connected
}
