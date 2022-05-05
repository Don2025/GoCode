package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

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

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		n, _ := strconv.Atoi(input.Text())
		grid1 := make([][]int, n)
		for i := 0; i < n; i++ {
			input.Scan()
			grid1[i] = stringArrayToIntArray(strings.Fields(input.Text()))
		}
		grid2 := make([][]int, n)
		for i := 0; i < n; i++ {
			input.Scan()
			grid2[i] = stringArrayToIntArray(strings.Fields(input.Text()))
		}
		println(countSubIslands(grid1, grid2))
	}
}

func stringArrayToIntArray(strings []string) []int {
	var arr []int
	for _, x := range strings {
		n, _ := strconv.Atoi(x)
		arr = append(arr, n)
	}
	return arr
}

/*
 * 执行用时：252ms 在所有Go提交中击败了84.66%的用户
 * 占用内存：21MB 在所有Go提交中击败了31.29%的用户
**/
