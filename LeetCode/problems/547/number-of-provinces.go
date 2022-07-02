package main

import (
	"bufio"
	. "fmt"
	"github.com/Don2025/GoCode/utils"
	"os"
	"strconv"
)

// https://leetcode.cn/problems/number-of-provinces/
//------------------------Leetcode Problem 547------------------------
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

//------------------------Leetcode Problem 547------------------------
/*
 * https://leetcode.cn/problems/number-of-provinces/
 * 执行用时：16ms 在所有Go提交中击败了96.19%的用户
 * 占用内存：6.4MB 在所有Go提交中击败了99.83%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		n, _ := strconv.Atoi(scanner.Text())
		isConnected := make([][]int, n)
		for i := 0; i < n; i++ {
			isConnected[i] = utils.StringToInts(scanner.Text())
		}
		Printf("Output: %v\n", findCircleNum(isConnected))
	}
}
