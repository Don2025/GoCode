package main

import (
	"bufio"
	. "fmt"
	"github.com/Don2025/GoCode/utils"
	"os"
	"strconv"
)

// https://leetcode.cn/problems/the-time-when-the-network-becomes-idle/
//------------------------Leetcode Problem 2039------------------------
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

/*
 * https://leetcode.cn/problems/the-time-when-the-network-becomes-idle/
 * 执行用时：284ms 在所有Go提交中击败了88.89%的用户
 * 占用内存：28.8MB 在所有Go提交中击败了44.44%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		n, _ := strconv.Atoi(scanner.Text())
		edges := make([][]int, n)
		for i := 0; i < n; i++ {
			scanner.Scan()
			edges[i] = utils.StringToInts(scanner.Text())
		}
		scanner.Scan()
		patience := utils.StringToInts(scanner.Text())
		Printf("Output: %d\n", networkBecomesIdle(edges, patience))
	}
}
