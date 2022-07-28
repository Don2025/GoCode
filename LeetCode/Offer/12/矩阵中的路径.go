package main

import (
	"bufio"
	. "fmt"
	"os"
	"strconv"
)

// https://leetcode.cn/problems/ju-zhen-zhong-de-lu-jing-lcof/
// ------------------------剑指 Offer I Problem 12------------------------
func exist(board [][]byte, word string) bool {
	row, col := len(board), len(board[0])
	var dfs func(int, int, int) bool
	dfs = func(i int, j int, k int) bool {
		if i < 0 || i >= row || j < 0 || j >= col || board[i][j] != word[k] {
			return false
		}
		if k == len(word)-1 {
			return true
		}
		board[i][j] = '#'
		ans := dfs(i+1, j, k+1) || dfs(i-1, j, k+1) || dfs(i, j+1, k+1) || dfs(i, j-1, k+1)
		board[i][j] = word[k]
		return ans
	}
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if dfs(i, j, 0) {
				return true
			}
		}
	}
	return false
}

// ------------------------剑指 Offer I Problem 12------------------------
/*
 * https://leetcode.cn/problems/ju-zhen-zhong-de-lu-jing-lcof/
 * 执行用时：4ms 在所有Go提交中击败了89.97%的用户
 * 占用内存：3.1MB 在所有Go提交中击败了99.40%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		n, _ := strconv.Atoi(scanner.Text())
		board := make([][]byte, n)
		for i := range board {
			scanner.Scan()
			board[i] = []byte(scanner.Text())
		}
		scanner.Scan()
		word := scanner.Text()
		Printf("Output: %v\n", exist(board, word))
	}
}
