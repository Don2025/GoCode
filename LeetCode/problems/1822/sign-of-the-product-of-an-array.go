package main

import (
	"bufio"
	. "fmt"
	"github.com/Don2025/GoCode/utils"
	"os"
)

// https://leetcode.cn/problems/sign-of-the-product-of-an-array/
//------------------------Leetcode Problem 1822------------------------
func arraySign(nums []int) int {
	ans := 1
	for _, x := range nums {
		if x < 0 {
			ans = -ans
		} else if x == 0 {
			return 0
		}
	}
	return ans
}

//------------------------Leetcode Problem 1822------------------------
/*
 * https://leetcode.cn/problems/sign-of-the-product-of-an-array/
 * 执行用时：4ms 在所有Go提交中击败了95.28%的用户
 * 占用内存：3.1MB 在所有Go提交中击败了73.23%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		nums := utils.StringToInts(scanner.Text())
		Printf("Output: %v\n", arraySign(nums))
	}
}
