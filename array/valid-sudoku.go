func isValidSudokuCell(nums []byte) bool {
    seen := make([]bool, 10)
    for _, v1 := range(nums) {
        if v1 != '.' {
            v := int(v1-'0')
            if seen[v] == false {
                seen[v] = true
            } else {
                return false
            }
        } 
    }
    return true
}

func isValidSudoku(board [][]byte) bool {
    to_check := make([]byte, 9)
    
    for row := 0; row < 9; row++ {
        for col := 0; col < 9; col++ {
            //to_check = append(to_check, board[row][col])
            to_check[col] = board[row][col]
        }
        // check the row
        if ! isValidSudokuCell(to_check) {
            return false
        }
    }
    
    for col := 0; col < 9; col++ {
        for row := 0; row < 9; row++ {
            //to_check = append(to_check, board[row][col])
            to_check[row] = board[row][col]
        }
        // check the column
        if ! isValidSudokuCell(to_check) {
            return false
        }
    }
    
    for sr := 0; sr < 3; sr++ {
        for sc := 0; sc < 3; sc ++ {
            for cellid := 0; cellid < 9; cellid ++ {
                row := sr*3 + cellid / 3
                col := sc*3 + cellid % 3
                //to_check = append(to_check, board[row][col])
                to_check[cellid] = board[row][col]
            }
            // check the subboard
            if ! isValidSudokuCell(to_check) {
                return false
            }
        }
    }
    
    return true
}