package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

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

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		n, _ := strconv.Atoi(input.Text())
		graph := make([][]int, n)
		for i := 0; i < n; i++ {
			input.Scan()
			graph[i] = stringArrayToIntArray(strings.Fields(input.Text()))
		}
		println(shortestPathLength(graph))
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
 * 执行用时：8ms 在所有Go提交中击败了48.08%的用户
 * 占用内存：6.7MB 在所有Go提交中击败了19.23%的用户
**/
