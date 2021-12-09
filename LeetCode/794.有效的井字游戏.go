package main

import (
	"bufio"
	"os"
	"strings"
)

func judgeWin(board []string, c byte) bool {
	for i := 0; i < 3; i++ {
		if board[i][0] == c && board[i][1] == c && board[i][2] == c || board[0][i] == c && board[1][i] == c && board[2][i] == c {
			return true
		}
	}
	return board[0][0] == c && board[1][1] == c && board[2][2] == c || board[0][2] == c && board[1][1] == c && board[2][0] == c
}

func validTicTacToe(board []string) bool {
	var oCnt, xCnt int
	for _, row := range board {
		oCnt += strings.Count(row, "O")
		xCnt += strings.Count(row, "X")
	}
	return !(oCnt != xCnt && oCnt != xCnt-1 || oCnt != xCnt && judgeWin(board, 'O') || oCnt != xCnt-1 && judgeWin(board, 'X'))
}

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		board := make([]string, 3)
		board[0] = input.Text()
		for i := 1; i <= 2; i++ {
			input.Scan()
			board[i] = input.Text()
		}
		println(validTicTacToe(board))
	}
}

/*
 * 执行用时：0ms 在所有Go提交中击败了100.00%的用户
 * 占用内存：1.9MB 在所有Go提交中击败了56.25%的用户
**/
