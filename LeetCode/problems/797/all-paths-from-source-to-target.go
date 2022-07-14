package main

import (
	"bufio"
	. "fmt"
	"github.com/Don2025/GoCode/utils"
	"os"
	"strconv"
)

// https://leetcode.cn/problems/all-paths-from-source-to-target/
//------------------------Leetcode Problem 797------------------------
func allPathsSourceTarget(graph [][]int) (ans [][]int) {
	stk := []int{0}
	var dfs func(int)
	dfs = func(x int) {
		if x == len(graph)-1 {
			ans = append(ans, append([]int(nil), stk...))
			return
		}
		for _, y := range graph[x] {
			stk = append(stk, y)
			dfs(y)
			stk = stk[:len(stk)-1]
		}
	}
	dfs(0)
	return
}

//------------------------Leetcode Problem 797------------------------
/*
 * https://leetcode.cn/problems/all-paths-from-source-to-target/
 * 执行用时：8ms 在所有Go提交中击败了86.11%的用户
 * 占用内存：6.4MB 在所有Go提交中击败了96.76%的用户
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
		Printf("Output: %v\n", allPathsSourceTarget(graph))
	}
}
