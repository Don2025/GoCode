package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func maximalNetworkRank(n int, roads [][]int) int {
	g := make([][]int, n)
	for i := range g {
		g[i] = make([]int, n)
	}
	degree := make([]int, n)
	for _, road := range roads {
		g[road[0]][road[1]] = 1
		g[road[1]][road[0]] = 1
		degree[road[0]]++
		degree[road[1]]++
	}
	var ans int
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			cur := degree[i] + degree[j]
			if g[i][j] == 1 || g[j][i] == 1 {
				cur--
			}
			ans = max(ans, cur)
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
		input.Scan()
		m, _ := strconv.Atoi(input.Text())
		roads := make([][]int, m)
		for i := range roads {
			input.Scan()
			roads[i] = stringArrayToIntArray(strings.Fields(input.Text()))
		}
		println(maximalNetworkRank(n, roads))
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
 * 执行用时：40ms 在所有Go提交中击败了43.75%的用户
 * 占用内存：6.9MB 在所有Go提交中击败了81.25%的用户
**/
