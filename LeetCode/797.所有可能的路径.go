package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

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

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		n, _ := strconv.Atoi(input.Text())
		graph := make([]int, n)
		for i := 0; i < n; i++ {
			input.Scan()
			graph[i] = stringArrayToIntArray(strings.Fields(input.Text()))
		}
		ans := allPathsSourceTarget(graph)
		for _, v := range ans {
			for _, x := range v {
				fmt.Printf("%d ", x)
			}
			fmt.Println()
		}
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
 * 执行用时：8ms 在所有Go提交中击败了86.11%的用户
 * 占用内存：6.4MB 在所有Go提交中击败了96.76%的用户
**/
