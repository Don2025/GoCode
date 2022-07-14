package main

import (
	"bufio"
	. "fmt"
	"github.com/Don2025/GoCode/utils"
	"os"
	"strconv"
)

// https://leetcode.cn/problems/subarray-product-less-than-k/
//------------------------Leetcode Problem 713------------------------

func numSubarrayProductLessThanK(nums []int, k int) int {
	if k == 0 || k == 1 {
		return 0
	}
	prod, ans := 1, 0 // prod存储nums[l]~nums[r]的累积
	for l, r := 0, 0; r < len(nums); r++ {
		prod *= nums[r]
		for prod >= k {
			prod /= nums[l]
			l++
		}
		ans += r - l + 1
	}
	return ans
}

//------------------------Leetcode Problem 713------------------------
/*
 * https://leetcode.cn/problems/subarray-product-less-than-k/
 * 执行用时：76ms 在所有Go提交中击败了47.90%的用户
 * 占用内存：7.2MB 在所有Go提交中击败了37.89%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		nums := utils.StringToInts(scanner.Text())
		scanner.Scan()
		k, _ := strconv.Atoi(scanner.Text())
		Printf("Output: %v\n", numSubarrayProductLessThanK(nums, k))
	}
}
