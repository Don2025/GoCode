package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func findCircleNum(isConnected [][]int) int {
	visited := make([]bool, 1005)
	var dfs func(int)
	dfs = func(i int) {
		for j := range isConnected {
			if isConnected[i][j] != 0 && !visited[j] {
				visited[j] = true
				dfs(j)
			}
		}
	}
	var cnt int
	for i := range isConnected {
		if !visited[i] {
			dfs(i)
			cnt++
		}
	}
	return cnt
}

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		n, _ := strconv.Atoi(input.Text())
		isConnected := make([][]int, n)
		for i := 0; i < n; i++ {
			isConnected[i] = stringArrayToIntArray(strings.Fields(input.Text()))
		}
		println(findCircleNum(isConnected))
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
 * 执行用时：16ms 在所有Go提交中击败了96.19%的用户
 * 占用内存：6.4MB 在所有Go提交中击败了99.83%的用户
**/
