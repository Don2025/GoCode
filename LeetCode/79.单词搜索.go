package main

import (
	"bufio"
	"os"
	"strconv"
)

func exist(board [][]byte, word string) bool {
	row, col := len(board), len(board[0])
	visited := make([][]bool, row)
	for i := range visited {
		visited[i] = make([]bool, col)
	}
	var check func(int, int, int) bool
	check = func(i int, j int, k int) bool {
		if board[i][j] != word[k] {
			return false
		} else if k == len(word)-1 {
			return true
		}
		visited[i][j] = true
		var dirs = []struct{ x, y int }{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} // 上下左右
		var ans bool
		for _, dir := range dirs {
			newX, newY := i+dir.x, j+dir.y
			if newX >= 0 && newX < row && newY >= 0 && newY < col {
				if !visited[newX][newY] {
					if check(newX, newY, k+1) {
						ans = true
						break
					}
				}
			}
		}
		visited[i][j] = false
		return ans
	}
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if check(i, j, 0) {
				return true
			}
		}
	}
	return false
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
		input.Scan()
		word := input.Text()
		println(exist(board, word))
	}
}

/*
 * 执行用时：88ms 在所有Go提交中击败了60.11%的用户
 * 占用内存：1.9MB 在所有Go提交中击败了60.46%的用户
**/
