package main

import (
	"bufio"
	. "fmt"
	"github.com/Don2025/GoCode/utils"
	"os"
)

// https://leetcode.cn/problems/n-repeated-element-in-size-2n-array/
//------------------------Leetcode Problem 961------------------------
func repeatedNTimes(nums []int) int {
	m := make(map[int]bool)
	for _, x := range nums {
		if m[x] {
			return x
		}
		m[x] = true
	}
	return -1
}

//------------------------Leetcode Problem 961------------------------
/*
 * https://leetcode.cn/problems/n-repeated-element-in-size-2n-array/
 * 执行用时：20ms 在所有Go提交中击败了96.60%的用户
 * 占用内存：6.5MB 在所有Go提交中击败了75.90%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		nums := utils.StringToInts(scanner.Text())
		Printf("Output: %v\n", repeatedNTimes(nums))
	}
}
