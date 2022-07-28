package main

import (
	"bufio"
	. "fmt"
	"github.com/Don2025/GoCode/utils"
	"math"
	"os"
	"strconv"
)

// https://leetcode.cn/problems/2VG8Kg/
// ------------------------剑指 Offer II Problem 8------------------------
func minSubArrayLen(target int, nums []int) int {
	ans := math.MaxInt32
	var sum int
	for i, j := 0, 0; j < len(nums); j++ {
		sum += nums[j]
		for i <= j && sum >= target {
			ans = min(ans, j-i+1)
			sum -= nums[i]
			i++
		}
	}
	if ans == math.MaxInt32 {
		ans = 0
	}
	return ans
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		target, _ := strconv.Atoi(scanner.Text())
		scanner.Scan()
		nums := utils.StringToInts(scanner.Text())
		Printf("Output: %v\n", minSubArrayLen(target, nums))
	}
}

// ------------------------剑指 Offer II Problem 8------------------------
/*
 * https://leetcode.cn/problems/2VG8Kg/
 * 执行用时：4ms 在所有Go提交中击败了95.41%的用户
 * 占用内存：3.6MB 在所有Go提交中击败了70.64%的用户
**/
