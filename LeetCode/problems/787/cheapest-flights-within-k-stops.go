package main

import (
	"bufio"
	. "fmt"
	"github.com/Don2025/GoCode/utils"
	"os"
	"strconv"
)

// https://leetcode.cn/problems/cheapest-flights-within-k-stops/
//------------------------Leetcode Problem 787------------------------
func findCheapestPrice(n int, flights [][]int, src int, dst int, k int) int {
	const INF = 0x3f3f3f3f
	f := make([]int, n)
	for i := range f {
		f[i] = INF
	}
	f[src] = 0
	ans := INF
	for t := 1; t <= k+1; t++ {
		g := make([]int, n)
		for i := range g {
			g[i] = INF
		}
		for _, flight := range flights {
			j, i, cost := flight[0], flight[1], flight[2]
			g[i] = min(g[i], f[j]+cost)
		}
		f = g
		ans = min(ans, f[dst])
	}
	if ans == INF {
		ans = -1
	}
	return ans
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

//------------------------Leetcode Problem 787------------------------
/*
 * https://leetcode.cn/problems/cheapest-flights-within-k-stops/
 * 执行用时：8ms 在所有Go提交中击败了97.28%的用户
 * 占用内存：4.7MB 在所有Go提交中击败了89.80%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		n, _ := strconv.Atoi(scanner.Text())
		scanner.Scan()
		m, _ := strconv.Atoi(scanner.Text())
		flights := make([][]int, m)
		for i := range flights {
			scanner.Scan()
			flights[i] = utils.StringToInts(scanner.Text())
		}
		scanner.Scan()
		var src, dst, k int
		Sscanf(scanner.Text(), "%d %d %d", &src, &dst, &k)
		Printf("Output: %v\n", findCheapestPrice(n, flights, src, dst, k))
	}
}
