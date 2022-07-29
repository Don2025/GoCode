package main

import (
	"bufio"
	. "fmt"
	"github.com/Don2025/GoCode/utils"
	"os"
)

// https://leetcode.cn/problems/TVdhkn/
//-------------------------剑指 Offer II Problem 79------------------------
func subsets(nums []int) [][]int {
	var ans [][]int
	n := len(nums)
	for mask := 0; mask < 1<<n; mask++ {
		var a []int
		for i, x := range nums {
			if mask>>i&1 > 0 {
				a = append(a, x)
			}
		}
		ans = append(ans, append([]int(nil), a...))
	}
	return ans
}

//-------------------------剑指 Offer II Problem 79------------------------
/*
 * https://leetcode.cn/problems/TVdhkn/
 * 执行用时：0ms 在所有Go提交中击败了100.00%的用户
 * 占用内存：2.2MB 在所有Go提交中击败了9.78%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		nums := utils.StringToInts(scanner.Text())
		Printf("Output: %v\n", subsets(nums))
	}
}
