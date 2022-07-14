package main

import (
	"bufio"
	. "fmt"
	"github.com/Don2025/GoCode/utils"
	"os"
	"strconv"
)

// https://leetcode.cn/problems/shortest-path-visiting-all-nodes/
//------------------------Leetcode Problem 847------------------------
func shortestPathLength(graph [][]int) int {
	n := len(graph)
	var queue [][]int
	visited := make([][]bool, n)
	for i := 0; i < n; i++ {
		queue = append(queue, []int{i, 0, 1 << i})
		visited[i] = make([]bool, 1<<n)
		visited[i][1<<i] = true
	}
	for len(queue) != 0 {
		t := queue[0]
		queue = queue[1:]
		if t[2] == (1<<n)-1 {
			return t[1]
		}
		for i := 0; i < len(graph[t[0]]); i++ {
			v := graph[t[0]][i]
			s := t[2] | 1<<v
			if !visited[v][s] {
				queue = append(queue, []int{v, t[1] + 1, s})
				visited[v][s] = true
			}
		}
	}
	return 0
}

//------------------------Leetcode Problem 847------------------------
/*
 * https://leetcode.cn/problems/shortest-path-visiting-all-nodes/
 * 执行用时：8ms 在所有Go提交中击败了48.08%的用户
 * 占用内存：6.7MB 在所有Go提交中击败了19.23%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		n, _ := strconv.Atoi(scanner.Text())
		graph := make([][]int, n)
		for i := 0; i < n; i++ {
			scanner.Scan()
			graph[i] = utils.StringToInts(scanner.Text())
		}
		Printf("Output: %v\n", shortestPathLength(graph))
	}
}
