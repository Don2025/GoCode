package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func movingCount(m int, n int, k int) int {
	visited := make([][]bool, m)
	for i := range visited {
		visited[i] = make([]bool, n)
	}
	var dfs func(int, int) int
	dfs = func(i int, j int) int {
		if i < 0 || i >= m || j < 0 || j >= n || (i/10+i%10+j/10+j%10) > k || visited[i][j] {
			return 0
		}
		visited[i][j] = true
		return dfs(i+1, j) + dfs(i-1, j) + dfs(i, j+1) + dfs(i, j-1) + 1
	}
	return dfs(0, 0)
}

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		arr := strings.Fields(input.Text())
		m, _ := strconv.Atoi(arr[0])
		n, _ := strconv.Atoi(arr[1])
		k, _ := strconv.Atoi(arr[2])
		println(movingCount(m, n, k))
	}
}

/*
 * 执行用时：0ms 在所有Go提交中击败了100.00%的用户
 * 占用内存：2.4MB 在所有Go提交中击败了58.75%的用户
**/
