package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func canReach(arr []int, start int) bool {
	var dfs func([]int, int) bool
	dfs = func(arr []int, start int) bool {
		if start < 0 || start >= len(arr) || arr[start] == -1 {
			return false
		}
		step := arr[start]
		arr[start] = -1
		return step == 0 || dfs(arr, start+step) || dfs(arr, start-step)
	}
	return dfs(arr, start)
}

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		arr := stringArrayToIntArray(strings.Fields(input.Text()))
		input.Scan()
		start, _ := strconv.Atoi(input.Text())
		println(canReach(arr, start))
	}
}

/*
 * 执行用时：40ms 在所有Go提交中击败了57.89%的用户
 * 占用内存：10.6MB 在所有Go提交中击败了26.32%的用户
**/
