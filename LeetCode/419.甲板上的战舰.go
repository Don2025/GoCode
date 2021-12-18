package main

import (
	"bufio"
	"os"
	"strconv"
)

func countBattleships(board [][]byte) int {
	var cnt int
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if board[i][j] == 'X' && (i == 0 || board[i-1][j] == '.') && (j == 0 || board[i][j-1] == '.') {
				cnt++
			}
		}
	}
	return cnt
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
		println(countBattleships(board))
	}
}

/*
 * 执行用时：0ms 在所有Go提交中击败了100.00%的用户
 * 占用内存：2.3MB 在所有Go提交中击败了100.00%的用户
**/
