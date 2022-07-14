package main

import (
	"bufio"
	. "fmt"
	"os"
	"strings"
)

// https://leetcode.cn/problems/valid-tic-tac-toe-state/
//------------------------Leetcode Problem 994------------------------
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

//------------------------Leetcode Problem 994------------------------
/*
 * https://leetcode.cn/problems/valid-tic-tac-toe-state/
 * 执行用时：0ms 在所有Go提交中击败了100.00%的用户
 * 占用内存：1.9MB 在所有Go提交中击败了56.25%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		board := make([]string, 3)
		board[0] = scanner.Text()
		for i := 1; i <= 2; i++ {
			scanner.Scan()
			board[i] = scanner.Text()
		}
		Printf("Output: %v\n", validTicTacToe(board))
	}
}
