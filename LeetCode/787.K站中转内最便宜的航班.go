package main

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

/*
 * 执行用时：8ms 在所有Go提交中击败了97.28%的用户
 * 占用内存：4.7MB 在所有Go提交中击败了89.80%的用户
**/
