package main

import (
	"bufio"
	. "fmt"
	"github.com/Don2025/GoCode/utils"
	"os"
	"strconv"
)

// https://leetcode.cn/problems/jump-game-iii/
//------------------------Leetcode Problem 1306------------------------
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

/*
 * https://leetcode.cn/problems/jump-game-iii/
 * 执行用时：40ms 在所有Go提交中击败了57.89%的用户
 * 占用内存：10.6MB 在所有Go提交中击败了26.32%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		arr := utils.StringToInts(scanner.Text())
		scanner.Scan()
		start, _ := strconv.Atoi(scanner.Text())
		Printf("Output: %v\n", canReach(arr, start))
	}
}
