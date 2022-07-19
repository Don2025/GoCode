package main

import (
	"bufio"
	. "fmt"
	"github.com/Don2025/GoCode/utils"
	"os"
	"strconv"
)

// https://leetcode.cn/problems/smallest-range-i/
//------------------------Leetcode Problem 908------------------------
func smallestRangeI(nums []int, k int) int {
	minVal, maxVal := nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] < minVal {
			minVal = nums[i]
		} else if nums[i] > maxVal {
			maxVal = nums[i]
		}
	}
	ans := maxVal - minVal - (k << 1)
	if ans > 0 {
		return ans
	}
	return 0
}

//------------------------Leetcode Problem 908------------------------
/*
 * https://leetcode.cn/problems/smallest-range-i/
 * 执行用时：12ms 在所有Go提交中击败了84.21%的用户
 * 占用内存：5.9MB 在所有Go提交中击败了92.98%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		nums := utils.StringToInts(scanner.Text())
		scanner.Scan()
		k, _ := strconv.Atoi(scanner.Text())
		Printf("Output: %v\n", smallestRangeI(nums, k))
	}
}
