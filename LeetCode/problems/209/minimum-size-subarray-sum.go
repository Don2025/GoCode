package main

import (
	"bufio"
	. "fmt"
	"github.com/Don2025/GoCode/utils"
	"os"
	"strconv"
)

// https://leetcode.cn/problems/minimum-size-subarray-sum/
//------------------------Leetcode Problem 209------------------------
func minSubArrayLen(target int, nums []int) int {
	var i, sum, ans int
	for j := 0; j < len(nums); j++ {
		sum += nums[j]
		for sum >= target {
			if ans == 0 {
				ans = j - i + 1
			} else {
				ans = min(ans, j-i+1)
			}
			sum -= nums[i]
			i++
		}
	}
	return ans
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

//------------------------Leetcode Problem 209------------------------
/*
 * https://leetcode.cn/problems/minimum-size-subarray-sum/
 * 执行用时：4ms 在所有Go提交中击败了96.58%的用户
 * 占用内存：3.6MB 在所有Go提交中击败了100.00%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		target, _ := strconv.Atoi(scanner.Text())
		scanner.Scan()
		nums := utils.StringToInts(scanner.Text())
		Printf("Output: %v\n", minSubArrayLen(target, nums))
	}
}
