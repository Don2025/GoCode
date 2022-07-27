package main

import (
	"bufio"
	. "fmt"
	"github.com/Don2025/GoCode/utils"
	"os"
	"strconv"
)

// https://leetcode.cn/problems/count-sub-islands/
//------------------------Leetcode Problem 1905------------------------
func countSubIslands(grid1 [][]int, grid2 [][]int) int {
	row, col := len(grid1), len(grid1[0])
	var flag bool
	var dfs func([][]int, [][]int, int, int)
	dfs = func(grid1 [][]int, grid2 [][]int, x int, y int) {
		if x < 0 || y < 0 || x >= row || y >= col || grid2[x][y] == 0 {
			return
		}
		grid2[x][y] = 0
		if grid1[x][y] == 0 {
			flag = false
		}
		dfs(grid1, grid2, x-1, y)
		dfs(grid1, grid2, x+1, y)
		dfs(grid1, grid2, x, y-1)
		dfs(grid1, grid2, x, y+1)
	}
	var ans int
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if grid2[i][j] == 1 {
				flag = true
				dfs(grid1, grid2, i, j)
				if flag {
					ans++
				}
			}
		}
	}
	return ans
}

//------------------------Leetcode Problem 1905------------------------
/*
 * https://leetcode.cn/problems/count-sub-islands/
 * 执行用时：252ms 在所有Go提交中击败了84.66%的用户
 * 占用内存：21MB 在所有Go提交中击败了31.29%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		n, _ := strconv.Atoi(scanner.Text())
		grid1 := make([][]int, n)
		for i := 0; i < n; i++ {
			scanner.Scan()
			grid1[i] = utils.StringToInts(scanner.Text())
		}
		grid2 := make([][]int, n)
		for i := 0; i < n; i++ {
			scanner.Scan()
			grid2[i] = utils.StringToInts(scanner.Text())
		}
		Printf("Output: %d\n", countSubIslands(grid1, grid2))
	}
}
