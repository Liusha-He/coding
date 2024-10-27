

def countPaths(grid: Sequence):
    n = len(grid)
    m = len(grid[0])

    grid_paths = [[0] * m for _ in range(n)]

    for i in range(n):
        if grid[i][0] == 1:
            grid_paths[i][0] = 1
        else:
            break

    for i in range(m):
        if grid[0][i] == 1:
            grid_paths[0][i] = 1
        else:
            break

    for i in range(1, n):
        for j in range(1, m):
            if grid[i][j] == 1:
                grid_paths[i][j] = grid_paths[i][j-1] + grid_paths[i-1][j]

    for g in grid_paths:
        print(g)

    return grid_paths[n-1][m-1]


if __name__ == '__main__':
    print(countPaths([[1,1,1,1,1],
                      [1,0,1,1,1],
                      [0,1,1,1,1],
                      [1,1,1,1,1]]))
