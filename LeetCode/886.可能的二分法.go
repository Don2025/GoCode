package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

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

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		n, _ := strconv.Atoi(input.Text())
		input.Scan()
		m, _ := strconv.Atoi(input.Text())
		dislikes := make([][]int, m)
		for i := range dislikes {
			input.Scan()
			dislikes[i] = stringArrayToIntArray(strings.Fields(input.Text()))
		}
		println(possibleBipartition(n, dislikes))
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
 * 执行用时：96ms 在所有Go提交中击败了70.73%的用户
 * 占用内存：8MB 在所有Go提交中击败了36.58%的用户
**/
