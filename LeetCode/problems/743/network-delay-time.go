package main

import (
	"bufio"
	. "fmt"
	"github.com/Don2025/GoCode/utils"
	"math"
	"os"
	"strconv"
)

// https://leetcode.cn/problems/network-delay-time/
//------------------------Leetcode Problem 743------------------------

func networkDelayTime(times [][]int, n, k int) (ans int) {
	const inf = math.MaxInt64 / 2
	g := make([][]int, n)
	for i := range g {
		g[i] = make([]int, n)
		for j := range g[i] {
			g[i][j] = inf
		}
	}
	for _, t := range times {
		x, y := t[0]-1, t[1]-1
		g[x][y] = t[2]
	}
	dist := make([]int, n)
	for i := range dist {
		dist[i] = inf
	}
	dist[k-1] = 0
	used := make([]bool, n)
	for i := 0; i < n; i++ {
		x := -1
		for y, u := range used {
			if !u && (x == -1 || dist[y] < dist[x]) {
				x = y
			}
		}
		used[x] = true
		for y, time := range g[x] {
			dist[y] = min(dist[y], dist[x]+time)
		}
	}
	for _, d := range dist {
		if d == inf {
			return -1
		}
		ans = max(ans, d)
	}
	return
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		n, _ := strconv.Atoi(scanner.Text())
		times := make([][]int, n)
		for i := range times {
			scanner.Scan()
			times[i] = utils.StringToInts(scanner.Text())
		}
		scanner.Scan()
		n, _ = strconv.Atoi(scanner.Text())
		scanner.Scan()
		k, _ := strconv.Atoi(scanner.Text())
		Printf("Output: %v\n", networkDelayTime(times, n, k))
	}
}

//------------------------Leetcode Problem 743------------------------
/*
 * https://leetcode.cn/problems/network-delay-time/
 * 执行用时：0ms 在所有Go提交中击败了30.53%的用户
 * 占用内存：2.4MB 在所有Go提交中击败了81.85%的用户
**/
