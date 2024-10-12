package syntax_sugar

import kotlin.random.Random

sealed class Matrix2D {
    abstract val rowCount: Int
    abstract val columnCount: Int
    abstract operator fun get(row: Int, col: Int): Int

    class ArrayMatrix(private val data: Array<IntArray>) : Matrix2D() {
        override val rowCount: Int get() = data.size
        override val columnCount: Int get() = if (data.isEmpty()) 0 else data[0].size
        override operator fun get(row: Int, col: Int): Int = data[row][col]
    }

    class ListMatrix(private val data: List<List<Int>>) : Matrix2D() {
        override val rowCount: Int get() = data.size
        override val columnCount: Int get() = if (data.isEmpty()) 0 else data[0].size
        override operator fun get(row: Int, col: Int): Int = data[row][col]
    }
}

fun findPaths(grid: Matrix2D): Int {
    val n = grid.rowCount
    val m = grid.columnCount

    val gridPaths = MutableList(n) {MutableList(m) {0} }

    for (i in 0..<n) {
        if (grid[i, 0] == 1) {
            gridPaths[i][0] = 1
        } else {
            break
        }
    }

    for (i in 0..<m) {
        if (grid[0, i] == 1) {
            gridPaths[0][i] = 1
        } else {
            break
        }
    }

    for (i in 1..<n) {
        for (j in 1..<m) {
            if (grid[i, j] == 0) {
                gridPaths[i][j] = 0
            } else {
                gridPaths[i][j] = gridPaths[i-1][j] + gridPaths[i][j-1]
            }
        }
    }
    return gridPaths[n-1][m-1]
}

fun main() {
    val nrow = 4
    val ncol = 4

    val grid = Array(nrow) { _ ->
        IntArray(ncol) { _ ->
            val r = Random.nextDouble()

            if (r < 0.1) {
                0
            } else {
                1
            }
        }
    }
    val gridList = List(nrow) { _ ->
        List(ncol) { _ ->
            val r = Random.nextDouble()

            if (r < 0.1) {
                0
            } else {
                1
            }
        }
    }

    grid[0][0] = 1
    grid[nrow - 1][ncol - 1] = 1
    grid.forEach { row -> println(row.toList()) }
    println(findPaths(Matrix2D.ArrayMatrix(grid)))

    println("==========")
    gridList.forEach { row -> println(row) }
    println(findPaths(Matrix2D.ListMatrix(gridList)))
}
