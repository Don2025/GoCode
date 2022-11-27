package main

import (
	"bufio"
	. "fmt"
	"github.com/Don2025/GoCode/utils"
	"os"
)

// https://leetcode.cn/problems/check-if-array-is-sorted-and-rotated/
//------------------------Leetcode Problem 1752------------------------
func check(nums []int) bool {
	cnt := 0
	for i := range nums {
		if nums[i] > nums[(i+1)%len(nums)] {
			cnt++
		}
	}
	return cnt <= 1
}

//------------------------Leetcode Problem 1752------------------------
/*
 * https://leetcode.cn/problems/check-if-array-is-sorted-and-rotated/
 * 执行用时：0ms 在所有Go提交中击败了100.00%的用户
 * 占用内存：2MB 在所有Go提交中击败了50.00%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		nums := utils.StringToInts(scanner.Text())
		Printf("Output: %v\n", check(nums))
	}
}
