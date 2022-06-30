package main

import (
	"bufio"
	. "fmt"
	"os"
	"strconv"
)

// https://leetcode.cn/problems/battleships-in-a-board/
//------------------------Leetcode Problem 419------------------------
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

/*
 * https://leetcode.cn/problems/battleships-in-a-board/
 * 执行用时：0ms 在所有Go提交中击败了100.00%的用户
 * 占用内存：2.3MB 在所有Go提交中击败了100.00%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		n, _ := strconv.Atoi(scanner.Text())
		board := make([][]byte, n)
		for i := 0; i < n; i++ {
			scanner.Scan()
			board[i] = []byte(scanner.Text())
		}
		Printf("Output: %v\n", countBattleships(board))
	}
}
