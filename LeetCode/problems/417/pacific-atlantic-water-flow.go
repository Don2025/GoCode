package main

import (
	"bufio"
	. "fmt"
	"github.com/Don2025/GoCode/utils"
	"os"
	"strconv"
)

// https://leetcode.cn/problems/pacific-atlantic-water-flow/
//------------------------Leetcode Problem 417------------------------
func pacificAtlantic(heights [][]int) [][]int {
	var ans [][]int
	m := len(heights)
	if m < 1 {
		return ans
	}
	n := len(heights[0])
	Pacific := make([][]bool, m)
	Atlantic := make([][]bool, m)
	for i := 0; i < m; i++ {
		Pacific[i] = make([]bool, n)
		Atlantic[i] = make([]bool, n)
	}
	var dfs func(int, int, [][]bool, int)
	dfs = func(x int, y int, visited [][]bool, pre int) {
		if x < 0 || y < 0 || x >= m || y >= n || visited[x][y] || heights[x][y] < pre {
			return
		}
		visited[x][y] = true
		dfs(x+1, y, visited, heights[x][y])
		dfs(x-1, y, visited, heights[x][y])
		dfs(x, y+1, visited, heights[x][y])
		dfs(x, y-1, visited, heights[x][y])
	}
	for i := 0; i < m; i++ {
		dfs(i, 0, Pacific, heights[i][0])
		dfs(i, n-1, Atlantic, heights[i][n-1])
	}
	for i := 0; i < n; i++ {
		dfs(0, i, Pacific, heights[0][i])
		dfs(m-1, i, Atlantic, heights[m-1][i])
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if Pacific[i][j] && Atlantic[i][j] {
				ans = append(ans, []int{i, j})
			}
		}
	}
	return ans
}

//------------------------Leetcode Problem 417------------------------
/*
 * https://leetcode.cn/problems/pacific-atlantic-water-flow/
 * 执行用时：28ms 在所有Go提交中击败了92.20%的用户
 * 占用内存：6.7MB 在所有Go提交中击败了84.40%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		n, _ := strconv.Atoi(scanner.Text())
		heights := make([][]int, n)
		for i := 0; i < n; i++ {
			scanner.Scan()
			heights[i] = utils.StringToInts(scanner.Text())
		}
		Printf("Output: %v\n", pacificAtlantic(heights))
	}
}
