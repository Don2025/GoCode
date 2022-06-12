package main

import (
	"bufio"
	. "fmt"
	"github.com/Don2025/GoCode/utils"
	"os"
	"strconv"
)

// https://leetcode.cn/problems/bP4bmD/
// ------------------------剑指 Offer II Problem 110------------------------
func allPathsSourceTarget(graph [][]int) (ans [][]int) {
	stack := []int{0}
	var dfs func(int)
	dfs = func(x int) {
		if x == len(graph)-1 {
			ans = append(ans, append([]int(nil), stack...))
			return
		}
		for _, y := range graph[x] {
			stack = append(stack, y)
			dfs(y)
			stack = stack[:len(stack)-1]
		}
	}
	dfs(0)
	return
}

// ------------------------剑指 Offer II Problem 110------------------------
/*
 * https://leetcode.cn/problems/bP4bmD/
 * 执行用时：8ms 在所有Go提交中击败了79.13%的用户
 * 占用内存：6.4MB 在所有Go提交中击败了75.65%的用户
**/
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	Printf("Input the number of rows of matrix:")
	for scanner.Scan() {
		n, _ := strconv.Atoi(scanner.Text())
		graph := make([][]int, n)
		for i := range graph {
			Printf("Input the %d row of numbers separated by spaces:\n", i+1)
			scanner.Scan()
			graph[i] = utils.StringToInts(scanner.Text())
		}
		Printf("Output: %v\n", allPathsSourceTarget(graph))
		Printf("Input the number of rows of matrix:")
	}
}
