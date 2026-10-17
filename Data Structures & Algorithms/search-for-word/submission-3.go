func exist(board [][]byte, word string) bool {
    var dfs func(int, int, int) bool
    dfs = func(i, j, x int) bool {
        // 1. Success condition
        if x == len(word) {
            return true
        }
        
        // 2. Boundary checks (Must happen BEFORE accessing board[i][j])
        if i < 0 || j < 0 || i == len(board) || j == len(board[0]) {
            return false
        }
        
        // 3. Character match check
        if word[x] != board[i][j] {
            return false
        }
        
        // 4. Backtracking: Mark cell as visited so DFS doesn't reuse it
        temp := board[i][j]
        board[i][j] = '#' // Use a placeholder character
        
        // Explore all 4 neighbors
        found := dfs(i, j+1, x+1) || 
                 dfs(i+1, j, x+1) || 
                 dfs(i, j-1, x+1) || 
                 dfs(i-1, j, x+1)
        
        // 5. Restore the cell's original value for other paths
        board[i][j] = temp
        
        return found
    }
    
    for i := 0; i < len(board); i++ {
        for j := 0; j < len(board[i]); j++ {
            if dfs(i, j, 0) {
                return true
            }
        }
    }
    return false
}
