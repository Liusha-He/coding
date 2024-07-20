package syntax_sugar

fun findPaths(grid: List<List<Int>>): Int {
    val n = grid.size
    val m = grid[0].size

    val gridPaths = MutableList(n) {MutableList(m) {0} }

    for (i in 0..n-1) {
        if (grid[i][0] == 1) {
            gridPaths[i][0] = 1
        } else {
            break
        }
    }

    for (i in 0..m-1) {
        if (grid[0][i] == 1) {
            gridPaths[0][i] = 1
        } else {
            break
        }
    }

    for (i in 1..n-1) {
        for (j in 1..m-1) {
            if (grid[i][j] == 0) {
                gridPaths[i][j] = 0
            } else {
                gridPaths[i][j] = gridPaths[i-1][j] + gridPaths[i][j-1]
            }
        }
    }
    return gridPaths[n-1][m-1]
}

fun main() {
    val grid = listOf(
        listOf(1,1,1,1),
        listOf(1,0,1,1),
        listOf(1,1,1,1),
        listOf(1,1,1,1)
    )
    println(findPaths(grid))
}
