package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

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

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		n, _ := strconv.Atoi(input.Text())
		heights := make([][]int, n)
		for i := 0; i < n; i++ {
			input.Scan()
			heights[i] = stringArrayToIntArray(strings.Fields(input.Text()))
		}
		ans := pacificAtlantic(heights)
		for _, x := range ans {
			fmt.Printf("%v ", x)
		}
		fmt.Println()
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
 * 执行用时：28ms 在所有Go提交中击败了92.20%的用户
 * 占用内存：6.7MB 在所有Go提交中击败了84.40%的用户
**/
