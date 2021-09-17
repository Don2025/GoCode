package main

import (
	"bufio"
	"os"
	"strconv"
)

func isValidSudoku(board [][]byte) bool {
	var rows, columns [9][9]int
	var subboxes [3][3][9]int
	for i, row := range board {
		for j, c := range row {
			if c == '.' {
				continue
			}
			index := c - '1'
			rows[i][index]++
			columns[j][index]++
			subboxes[i/3][j/3][index]++
			if rows[i][index] > 1 || columns[j][index] > 1 || subboxes[i/3][j/3][index] > 1 {
				return false
			}
		}
	}
	return true
}

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		n, _ := strconv.Atoi(input.Text())
		board := make([][]byte, n)
		for i := 0; i < n; i++ {
			input.Scan()
			board[i] = []byte(input.Text())
		}
		println(isValidSudoku(board))
	}
}

/*
 * 执行用时：4ms 在所有Go提交中击败了67.61%的用户
 * 占用内存：2.6MB 在所有Go提交中击败了99.19%的用户
**/
