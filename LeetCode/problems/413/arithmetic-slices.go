package main

import (
	"bufio"
	. "fmt"
	"github.com/Don2025/GoCode/utils"
	"os"
)

// https://leetcode.cn/problems/arithmetic-slices/
//------------------------Leetcode Problem 413------------------------
func numberOfArithmeticSlices(nums []int) int {
	if len(nums) <= 2 {
		return 0
	}
	var ans int
	var add int
	for i := 2; i < len(nums); i++ {
		if nums[i-1]-nums[i] == nums[i-2]-nums[i-1] {
			add++
			ans += add
		} else {
			add = 0
		}
	}
	return ans
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		nums := utils.StringToInts(scanner.Text())
		Println(numberOfArithmeticSlices(nums))
	}
}

//------------------------Leetcode Problem 413------------------------
/*
 * https://leetcode.cn/problems/arithmetic-slices/
 * 执行用时：0ms 在所有Go提交中击败了100.00%的用户
 * 占用内存：2.2MB 在所有Go提交中击败了100.00%的用户
**/
