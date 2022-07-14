package main

import (
	"bufio"
	. "fmt"
	"github.com/Don2025/GoCode/utils"
	"os"
	"strconv"
)

// https://leetcode.cn/problems/is-graph-bipartite/
//------------------------Leetcode Problem 785------------------------
func isBipartite(graph [][]int) bool {
	UNCOLORED, RED, GREEN := 0, 1, 2
	color := make([]int, len(graph))
	valid := true
	var dfs func([][]int, int, int)
	dfs = func(g [][]int, x int, c int) {
		color[x] = c
		cNei := RED
		if c == RED {
			cNei = GREEN
		}
		for _, neighbor := range graph[x] {
			if color[neighbor] == UNCOLORED {
				dfs(graph, neighbor, cNei)
				if !valid {
					return
				}
			} else if color[neighbor] != cNei {
				valid = false
				return
			}
		}
	}
	for i := 0; i < len(graph) && valid; i++ {
		if color[i] == UNCOLORED {
			dfs(graph, i, RED)
		}
	}
	return valid
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		n, _ := strconv.Atoi(scanner.Text())
		graph := make([][]int, n)
		for i := range graph {
			scanner.Scan()
			graph[i] = utils.StringToInts(scanner.Text())
		}
		Printf("Output: %v\n", isBipartite(graph))
	}
}

//------------------------Leetcode Problem 785------------------------
/*
 * https://leetcode.cn/problems/is-graph-bipartite/
 * 执行用时：24ms 在所有Go提交中击败了39.10%的用户
 * 占用内存：6.5MB 在所有Go提交中击败了65.79%的用户
**/
