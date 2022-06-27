package main

import (
	"bufio"
	. "fmt"
	"github.com/Don2025/GoCode/utils"
	"os"
)

// https://leetcode.cn/problems/missing-number/
//------------------------Leetcode Problem 268------------------------
func missingNumber(nums []int) int {
	hash := make([]bool, len(nums)+1)
	for _, x := range nums {
		hash[x] = true
	}
	var ans int
	for i, x := range hash {
		if x == false {
			ans = i
		}
	}
	return ans
}

//------------------------Leetcode Problem 268------------------------
/*
 * https://leetcode.cn/problems/missing-number/
 * 执行用时：16ms 在所有Go提交中击败了86.81%的用户
 * 占用内存：6.2MB 在所有Go提交中击败了11.90%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		nums := utils.StringToInts(scanner.Text())
		Printf("Output: %v\n", missingNumber(nums))
	}
}
