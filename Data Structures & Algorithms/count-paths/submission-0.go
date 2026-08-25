// start = grid[0][0]
// end = grid[m - 1][n - 1]

// meet is when start == end

// action at each position:
// right:
// 	- grid[m][n + 1]
// down:
// 	- grid[m + 1][n]

// so combinations[m][n] is how many combinations of down + right actions until start == end

// -> combinations[m][n] = combinationsWhenWalk(right) + combinationsWhenWalk(down)

// -> combinations[m][n] = combinationsWhenWalk(grid[m][n + 1]) + combinationsWhenWalk(grid[m + 1][n])

// recursion relation met

// Observe: 
// 	- m:
// 		calculate m + 1 before m -> big to small
// 	- n:
// 		calculate n + 1 before n -> big to small

// -> this recursive relation actual start at end and caculation combinations from end to start (which is ok)

// We can expirement mentally to flip this around, what if we actually walk from start to end?
// then m, n relations are small to big

// base case then would be: grid[0][0] = 0

// actions:
// 	down: grid[1][0] = 1 + grid[0][0] + 0 = 1
// 	right: grid[0][1] = 1 + grid[0][0] + 0 = 1

// down -> right:
// 	grid[1][1] = 1 + grid[1][0] + grid[0][1] = 3

// --->>>

// at position: look add value of from left + right + 1 (1 to count itself)
// grid[m][n] = 1 + grid[m-1][n] + grid[m][n-1] = 1

// we should create extra col and row so grid column 0 and grid column 1 could do this same logic

func uniquePaths(m int, n int) int {
    grid := make([][]int, m)
	for i:=0; i<len(grid); i++ {
		grid[i] = make([]int, n)
	}

	for i:=0; i<m; i++{
		grid[i][0] = 1
	}

	for i:=0; i<n; i++{
		grid[0][i] = 1
	}

	for i:=1; i<m; i++ {
		for j:=1; j<n; j++ {
			grid[i][j] = grid[i-1][j] + grid[i][j-1] 
		}
	}

	return grid[m-1][n-1]
}
