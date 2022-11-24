package main

import (
	"bufio"
	. "fmt"
	"github.com/Don2025/GoCode/utils"
	"os"
)

// https://leetcode.cn/problems/number-of-subarrays-with-bounded-maximum/
//------------------------Leetcode Problem 795------------------------
func numSubarrayBoundedMax(nums []int, left int, right int) int {
	count := func(x int) int {
		var cur, ans int
		for _, num := range nums {
			if num <= x {
				cur++
			} else {
				cur = 0
			}
			ans += cur
		}
		return ans
	}
	return count(right) - count(left-1)
}

//------------------------Leetcode Problem 795------------------------
/*
 * https://leetcode.cn/problems/number-of-subarrays-with-bounded-maximum/
 * 执行用时：56ms 在所有Go提交中击败了44.83%的用户
 * 占用内存：7.8MB 在所有Go提交中击败了55.17%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		nums := utils.StringToInts(scanner.Text())
		scanner.Scan()
		var left, right int
		Sscanf(scanner.Text(), "%d %d", &left, &right)
		Printf("Output: %v\n", numSubarrayBoundedMax(nums, left, right))
	}
}
