package main

import (
	"bufio"
	. "fmt"
	"github.com/Don2025/GoCode/utils"
	"os"
	"strconv"
)

// https://leetcode.cn/problems/bLyHh0/
// ------------------------剑指 Offer II Problem 116------------------------
func findCircleNum(isConnected [][]int) (ans int) {
	vis := make([]bool, len(isConnected))
	for i, v := range vis {
		if !v {
			ans++
			queue := []int{i}
			for len(queue) > 0 {
				from := queue[0]
				queue = queue[1:]
				vis[from] = true
				for to, conn := range isConnected[from] {
					if conn == 1 && !vis[to] {
						queue = append(queue, to)
					}
				}
			}
		}
	}
	return
}

// ------------------------剑指 Offer II Problem 116------------------------
/*
 * https://leetcode.cn/problems/bLyHh0/
 * 执行用时：24ms 在所有Go提交中击败了21.43%的用户
 * 占用内存：6.5MB 在所有Go提交中击败了17.14%的用户
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
		Printf("Output: %v\n", findCircleNum(graph))
		Printf("Input the number of rows of matrix:")
	}
}
