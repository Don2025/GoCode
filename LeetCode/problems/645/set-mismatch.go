package main

import (
	"bufio"
	. "fmt"
	"github.com/Don2025/GoCode/utils"
	"os"
)

// https://leetcode.cn/problems/set-mismatch/
//------------------------Leetcode Problem 645------------------------
func findErrorNums(nums []int) []int {
	mp := make([]int, len(nums)+1)
	for _, x := range nums {
		mp[x]++
	}
	ans := make([]int, 2)
	for i, x := range mp {
		if x > 1 {
			ans[0] = i
		} else if x == 0 {
			ans[1] = i
		}
	}
	return ans
}

//------------------------Leetcode Problem 645------------------------
/*
 * https://leetcode.cn/problems/set-mismatch/
 * 执行用时：32ms 在所有Go提交中击败了72.48%的用户
 * 占用内存：6.5MB 在所有Go提交中击败了53.69%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		nums := utils.StringToInts(scanner.Text())
		Printf("Output: %v\n", findErrorNums(nums))
	}
}
