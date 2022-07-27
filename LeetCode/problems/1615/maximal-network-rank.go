package main

import (
	"bufio"
	. "fmt"
	"github.com/Don2025/GoCode/utils"
	"os"
	"strconv"
)

// https://leetcode.cn/problems/maximal-network-rank/
//------------------------Leetcode Problem 1615------------------------
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

//------------------------Leetcode Problem 1615------------------------
/*
 * https://leetcode.cn/problems/maximal-network-rank/
 * 执行用时：40ms 在所有Go提交中击败了43.75%的用户
 * 占用内存：6.9MB 在所有Go提交中击败了81.25%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		n, _ := strconv.Atoi(scanner.Text())
		scanner.Scan()
		m, _ := strconv.Atoi(scanner.Text())
		roads := make([][]int, m)
		for i := range roads {
			scanner.Scan()
			roads[i] = utils.StringToInts(scanner.Text())
		}
		Printf("Output: %v\n", maximalNetworkRank(n, roads))
	}
}
