package coding

func FindPaths(grid [][]int) int {
	n := len(grid)
	m := len(grid[0])
	gridPaths := make([][]int, n)
	for i := range gridPaths {
		gridPaths[i] = make([]int, m)
	}

	for i := 0; i < n; i++ {
		if grid[i][0] == 1 {
			gridPaths[i][0] = 1
		} else {
			break
		}
	}

	for i := 0; i < m; i++ {
		if grid[0][i] == 1 {
			gridPaths[0][i] = 1
		} else {
			break
		}
	}

	for i := 1; i < n; i++ {
		for j := 1; j < m; j++ {
			if grid[i][j] == 0 {
				gridPaths[i][j] = 0
			} else {
				gridPaths[i][j] = gridPaths[i-1][j] + gridPaths[i][j-1]
			}
		}
	}
	return gridPaths[n-1][m-1]
}
