package main

import (
	"bufio"
	. "fmt"
	"os"
)

// https://leetcode.cn/problems/ji-qi-ren-de-yun-dong-fan-wei-lcof/
// ------------------------剑指 Offer I Problem 13------------------------
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

// ------------------------剑指 Offer I Problem 13------------------------
/*
 * https://leetcode.cn/problems/ji-qi-ren-de-yun-dong-fan-wei-lcof/
 * 执行用时：0ms 在所有Go提交中击败了100.00%的用户
 * 占用内存：2.4MB 在所有Go提交中击败了58.75%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		var m, n, k int
		Sscanf(scanner.Text(), "%d %d %d", &m, &n, &k)
		Printf("Output: %d\n", movingCount(m, n, k))
	}
}
