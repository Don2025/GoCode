package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func solve(board [][]byte) {
	row, col := len(board), len(board[0])
	if row == 0 || col == 0 {
		return
	}
	var dfs func(int, int)
	dfs = func(x int, y int) {
		if x < 0 || x >= row || y < 0 || y >= col || board[x][y] != 'O' {
			return
		}
		board[x][y] = 'A'
		dfs(x+1, y)
		dfs(x-1, y)
		dfs(x, y+1)
		dfs(x, y-1)
	}
	for i := 0; i < row; i++ {
		dfs(i, 0)
		dfs(i, col-1)
	}
	for i := 1; i < col-1; i++ {
		dfs(0, i)
		dfs(row-1, i)
	}
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if board[i][j] == 'A' {
				board[i][j] = 'O'
			} else if board[i][j] == 'O' {
				board[i][j] = 'X'
			}
		}
	}
}

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		n, _ := strconv.Atoi(input.Text())
		board := make([][]byte, n)
		for i := range board {
			input.Scan()
			board[i] = []byte(input.Text())
		}
		solve(board)
		for _, s := range board {
			fmt.Printf("%s\n", s)
		}
	}
}

/*
 * 执行用时：20ms 在所有Go提交中击败了44.75%的用户
 * 占用内存：6.1MB 在所有Go提交中击败了81.41%的用户
**/
