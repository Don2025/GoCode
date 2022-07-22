package main

import (
	"bufio"
	. "fmt"
	"github.com/Don2025/GoCode/utils"
	"os"
	"sort"
)

// https://leetcode.cn/problems/largest-perimeter-triangle/
//------------------------Leetcode Problem 976------------------------
func largestPerimeter(nums []int) int {
	sort.Ints(nums)
	for i := len(nums) - 1; i >= 2; i-- {
		if nums[i-2]+nums[i-1] > nums[i] {
			return nums[i-2] + nums[i-1] + nums[i]
		}
	}
	return 0
}

//------------------------Leetcode Problem 976------------------------
/*
 * https://leetcode.cn/problems/largest-perimeter-triangle/
 * 执行用时：56ms 在所有Go提交中击败了6.67%的用户
 * 占用内存：6.4MB 在所有Go提交中击败了48.57%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		nums := utils.StringToInts(scanner.Text())
		Printf("Output: %v\n", largestPerimeter(nums))
	}
}
