package main

import (
	"bufio"
	. "fmt"
	"github.com/Don2025/GoCode/utils"
	"os"
	"strconv"
)

// https://leetcode.cn/problems/ZVAVXX/
// ------------------------剑指 Offer II Problem 9------------------------
func numSubarrayProductLessThanK(nums []int, k int) int {
	if k <= 1 {
		return 0
	}
	var ans int
	mul := 1
	for l, r := 0, 0; r < len(nums); r++ {
		mul *= nums[r]
		for mul >= k {
			mul /= nums[l]
			l++
		}
		ans += r - l + 1
	}
	return ans
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		nums := utils.StringToInts(scanner.Text())
		scanner.Scan()
		k, _ := strconv.Atoi(scanner.Text())
		Printf("Output: %v\n", numSubarrayProductLessThanK(nums, k))
	}
}

// ------------------------剑指 Offer II Problem 9------------------------
/*
 * https://leetcode.cn/problems/ZVAVXX/
 * 执行用时：64ms 在所有Go提交中击败了95.08%的用户
 * 占用内存：7.2MB 在所有Go提交中击败了66.67%的用户
**/
