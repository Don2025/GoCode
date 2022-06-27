package main

import (
	"bufio"
	. "fmt"
	"github.com/Don2025/GoCode/utils"
	"os"
	"strconv"
)

// https://leetcode.cn/problems/minimum-height-trees/
//------------------------Leetcode Problem 310------------------------
func findMinHeightTrees(n int, edges [][]int) []int {
	if n == 1 {
		return []int{0}
	}
	g := make([][]int, n)
	for _, e := range edges {
		x, y := e[0], e[1]
		g[x] = append(g[x], y)
		g[y] = append(g[y], x)
	}
	parents := make([]int, n)
	maxDepth, node := 0, -1
	var dfs func(int, int, int)
	dfs = func(x, pa, depth int) {
		if depth > maxDepth {
			maxDepth, node = depth, x
		}
		parents[x] = pa
		for _, y := range g[x] {
			if y != pa {
				dfs(y, x, depth+1)
			}
		}
	}
	dfs(0, -1, 1)
	maxDepth = 0
	dfs(node, -1, 1)
	var path []int
	for node != -1 {
		path = append(path, node)
		node = parents[node]
	}
	m := len(path)
	if m%2 == 0 {
		return []int{path[m/2-1], path[m>>1]}
	}
	return []int{path[m>>1]}
}

//------------------------Leetcode Problem 310------------------------
/*
 * https://leetcode.cn/problems/minimum-height-trees/
 * 执行用时：60ms 在所有Go提交中击败了96.69%的用户
 * 占用内存：11.1MB 在所有Go提交中击败了15.71%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		n, _ := strconv.Atoi(scanner.Text())
		scanner.Scan()
		m, _ := strconv.Atoi(scanner.Text())
		edges := make([][]int, m)
		for i := range edges {
			scanner.Scan()
			edges[i] = utils.StringToInts(scanner.Text())
		}
		Printf("Output: %v\n", findMinHeightTrees(n, edges))
	}
}
