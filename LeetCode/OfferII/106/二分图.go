package main

import (
	"bufio"
	. "fmt"
	"github.com/Don2025/GoCode/utils"
	"os"
	"strconv"
)

// https://leetcode.cn/problems/vEAB3K/
// ------------------------剑指 Offer II Problem 106------------------------
func isBipartite(graph [][]int) bool {
	visited := make([]int, len(graph))
	var queue []int
	for i := 0; i < len(graph); i++ {
		if visited[i] != 0 {
			continue
		}
		queue = append(queue, i)
		visited[i] = 1
		for len(queue) != 0 {
			v := queue[0]
			queue = queue[1:]
			for _, w := range graph[v] {
				if visited[w] == 0 {
					visited[w] = -visited[v]
					queue = append(queue, w)
				}
				if visited[w] == visited[v] {
					return false
				}
			}
		}
	}
	return true
}

// ------------------------剑指 Offer II Problem 106------------------------
/*
 * https://leetcode.cn/problems/vEAB3K/
 * 执行用时：40ms 在所有Go提交中击败了7.02%的用户
 * 占用内存：6.4MB 在所有Go提交中击败了86.84%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	Printf("Input the number of rows of matrix:")
	for scanner.Scan() {
		n, _ := strconv.Atoi(scanner.Text())
		grid := make([][]int, n)
		for i := range grid {
			Printf("Input the %d row of numbers separated by spaces:\n", i+1)
			scanner.Scan()
			grid[i] = utils.StringToInts(scanner.Text())
		}
		Printf("Output: %v\n", isBipartite(grid))
	}
}
