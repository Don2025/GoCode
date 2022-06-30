package main

import (
	"bufio"
	. "fmt"
	"github.com/Don2025/GoCode/utils"
	"os"
)

// https://leetcode.cn/problems/arithmetic-slices-ii-subsequence/
// ------------------------Leetcode Problem 446------------------------
func numberOfArithmeticSlices(nums []int) int {
	var ans int
	mp := make([]map[int]int, len(nums))
	for i, x := range nums {
		mp[i] = map[int]int{}
		for j, y := range nums[:i] {
			diff := x - y
			cnt := mp[j][diff]
			ans += cnt
			mp[i][diff] += cnt + 1
		}
	}
	return ans
}

// ------------------------Leetcode Problem 446------------------------
/*
 * https://leetcode.cn/problems/arithmetic-slices-ii-subsequence/
 * 执行用时：176ms 在所有Go提交中击败了45.83%的用户
 * 占用内存：37.2MB 在所有Go提交中击败了58.33%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		nums := utils.StringToInts(scanner.Text())
		Printf("Output: %v\n", numberOfArithmeticSlices(nums))
	}
}
