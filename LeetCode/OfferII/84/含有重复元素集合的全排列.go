package main

import (
	"bufio"
	. "fmt"
	"github.com/Don2025/GoCode/utils"
	"os"
	"sort"
)

// https://leetcode.cn/problems/7p8L0Z/
// ------------------------剑指 Offer II Problem 84------------------------
func permuteUnique(nums []int) (ans [][]int) {
	sort.Ints(nums)
	n := len(nums)
	vis := make([]bool, n)
	var arr []int
	var backtrack func(int)
	backtrack = func(idx int) {
		if idx == n {
			ans = append(ans, append([]int(nil), arr...))
			return
		}
		for i, x := range nums {
			if vis[i] || i > 0 && !vis[i-1] && x == nums[i-1] {
				continue
			}
			arr = append(arr, x)
			vis[i] = true
			backtrack(idx + 1)
			vis[i] = false
			arr = arr[:len(arr)-1]
		}
	}
	backtrack(0)
	return
}

//------------------------剑指 Offer II Problem 84------------------------
/*
 * https://leetcode.cn/problems/7p8L0Z/
 * 执行用时：0ms 在所有Go提交中击败了100.00%的用户
 * 占用内存：3.6MB 在所有Go提交中击败了98.31%的用户
**/
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	Printf("Input a line of numbers separated by space:")
	for scanner.Scan() {
		Printf("Output: %v\n", permuteUnique(utils.StringToInts(scanner.Text())))
		Printf("Input a line of numbers separated by space:")
	}
}
