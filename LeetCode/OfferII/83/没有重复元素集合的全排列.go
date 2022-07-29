package main

import (
	"bufio"
	. "fmt"
	"github.com/Don2025/GoCode/utils"
	"os"
)

// https://leetcode.cn/problems/VvJkup/
//-------------------------剑指 Offer II Problem 83------------------------
func permute(nums []int) [][]int {
	var ans [][]int
	n := len(nums)
	vis := make([]bool, n)
	var path []int
	var dfs func()
	dfs = func() {
		if len(path) == n {
			ans = append(ans, append([]int(nil), path...))
			return
		}
		for i := 0; i < n; i++ {
			if !vis[i] {
				vis[i] = true
				path = append(path, nums[i])
				dfs()
				vis[i] = false
				path = path[:len(path)-1]
			}
		}
	}
	dfs()
	return ans
}

//-------------------------剑指 Offer II Problem 83------------------------
/*
 * https://leetcode.cn/problems/VvJkup/
 * 执行用时：0ms 在所有Go提交中击败了100.00%的用户
 * 占用内存：2.5MB 在所有Go提交中击败了79.87%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		nums := utils.StringToInts(scanner.Text())
		Printf("Output: %v\n", permute(nums))
	}
}
