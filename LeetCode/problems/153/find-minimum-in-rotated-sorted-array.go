package main

import (
	"bufio"
	. "fmt"
	"github.com/Don2025/GoCode/utils"
	"os"
	"sort"
)

// https://leetcode.cn/problems/find-minimum-in-rotated-sorted-array/
//------------------------Leetcode Problem 153------------------------
func findMin(nums []int) int {
	sort.Ints(nums)
	return nums[0]
}

//------------------------Leetcode Problem 153------------------------
/*
 * https://leetcode.cn/problems/find-minimum-in-rotated-sorted-array/
 * 执行用时：4ms 在所有Go提交中击败了11.73%的用户
 * 占用内存：2.3MB 在所有Go提交中击败了65.07%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		nums := utils.StringToInts(scanner.Text())
		Printf("Output: %v\n", findMin(nums))
	}
}
