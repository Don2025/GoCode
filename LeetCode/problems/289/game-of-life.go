package main

import (
	"bufio"
	. "fmt"
	"github.com/Don2025/GoCode/utils"
	"os"
	"strconv"
)

// https://leetcode.cn/problems/game-of-life/
//------------------------Leetcode Problem 289------------------------
func gameOfLife(board [][]int) {
	dirx := []int{-1, -1, -1, 0, 0, 1, 1, 1}
	diry := []int{-1, 0, 1, -1, 1, -1, 0, 1}
	for i := range board {
		for j := range board[i] {
			var cnt int
			for k := range dirx {
				if x, y := i+dirx[k], j+diry[k]; x >= 0 && x < len(board) && y >= 0 && y < len(board[i]) {
					cnt += board[x][y] & 1
				}
			}
			if cnt == 3 || cnt == 2 && board[i][j] == 1 {
				board[i][j] |= 2
			} else if cnt == 3 && board[i][j] == 0 {
				board[i][j] |= 2
			}
		}
	}
	for i := range board {
		for j := range board[i] {
			board[i][j] >>= 1
		}
	}
}

//------------------------Leetcode Problem 289------------------------
/*
 * https://leetcode.cn/problems/game-of-life/
 * 执行用时：0ms 在所有Go提交中击败了100.00%的用户
 * 占用内存：1.8MB 在所有Go提交中击败了87.35%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		n, _ := strconv.Atoi(scanner.Text())
		board := make([][]int, n)
		for i := range board {
			scanner.Scan()
			board[i] = utils.StringToInts(scanner.Text())
		}
		gameOfLife(board)
		Printf("Output: %v\n", board)
	}
}
