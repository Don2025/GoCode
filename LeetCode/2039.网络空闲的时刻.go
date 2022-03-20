package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func networkBecomesIdle(edges [][]int, patience []int) int {
	n := len(patience)
	g := make([][]int, n)
	for _, e := range edges {
		x, y := e[0], e[1]
		g[x] = append(g[x], y)
		g[y] = append(g[y], x)
	}
	visited := make([]bool, n)
	visited[0] = true
	queue := []int{0}
	var ans int
	for i := 1; queue != nil; i++ {
		q := queue
		queue = nil
		for _, x := range q {
			for _, v := range g[x] {
				if visited[v] {
					continue
				}
				visited[v] = true
				queue = append(queue, v)
				ans = max(ans, (i<<1-1)/patience[v]*patience[v]+i<<1+1)
			}
		}
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		n, _ := strconv.Atoi(input.Text())
		edges := make([][]int, n)
		for i := 0; i < n; i++ {
			input.Scan()
			edges[i] = stringArrayToIntArray(strings.Fields(input.Text()))
		}
		input.Scan()
		patience := stringArrayToIntArray(strings.Fields(input.Text()))
		println(networkBecomesIdle(edges, patience))
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
 * 执行用时：284ms 在所有Go提交中击败了88.89%的用户
 * 占用内存：28.8MB 在所有Go提交中击败了44.44%的用户
**/
