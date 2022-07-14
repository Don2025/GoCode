package main

import (
	"bufio"
	. "fmt"
	"github.com/Don2025/GoCode/utils"
	"os"
	"strconv"
)

// https://leetcode.cn/problems/possible-bipartition/
//------------------------Leetcode Problem 886------------------------
func possibleBipartition(n int, dislikes [][]int) bool {
	g := make([][]int, n+1)
	for _, d := range dislikes {
		g[d[0]] = append(g[d[0]], d[1])
		g[d[1]] = append(g[d[1]], d[0])
	}
	color := make(map[int]int)
	var dfs func(int, int) bool
	dfs = func(node int, c int) bool {
		if x, ok := color[node]; ok {
			return x == c
		}
		color[node] = c
		for _, next := range g[node] {
			if !dfs(next, c^1) {
				return false
			}
		}
		return true
	}
	for i := 1; i <= n; i++ {
		if _, ok := color[i]; !ok && !dfs(i, 0) {
			return false
		}
	}
	return true
}

//------------------------Leetcode Problem 886------------------------
/*
 * https://leetcode.cn/problems/possible-bipartition/
 * 执行用时：96ms 在所有Go提交中击败了70.73%的用户
 * 占用内存：8MB 在所有Go提交中击败了36.58%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		n, _ := strconv.Atoi(scanner.Text())
		scanner.Scan()
		m, _ := strconv.Atoi(scanner.Text())
		dislikes := make([][]int, m)
		for i := range dislikes {
			scanner.Scan()
			dislikes[i] = utils.StringToInts(scanner.Text())
		}
		Printf("Output: %v\n", possibleBipartition(n, dislikes))
	}
}
