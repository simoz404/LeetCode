package main

import "fmt"

func main() {
	b := [][]byte{
		{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
		{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
	}

	fmt.Println(isValidSudoku(b))
}

func isValidSudoku(board [][]byte) bool {
	var check = make(map[byte]bool)
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if check[board[i][j]] {
				check[board[i][j]] = false
			} else if board[i][j] != '.' {
				return false
			}
		}
		break
	}
	return true
}
