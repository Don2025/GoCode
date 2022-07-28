package main

import (
	"bufio"
	. "fmt"
	"github.com/Don2025/GoCode/utils"
	"os"
)

// https://leetcode.cn/problems/shu-zu-zhong-shu-zi-chu-xian-de-ci-shu-ii-lcof/
// ------------------------剑指 Offer I Problem 56-II------------------------
func singleNumber(nums []int) int {
	var a, b int
	for _, c := range nums {
		a, b = a&^c|b&c, b&^c|^a&^b&c
	}
	return ^a & b
}

// ------------------------剑指 Offer I Problem 56-II------------------------
/*
 * https://leetcode.cn/problems/shu-zu-zhong-shu-zi-chu-xian-de-ci-shu-ii-lcof/
 * 执行用时：16ms 在所有Go提交中击败了99.24%的用户
 * 占用内存：6.3MB 在所有Go提交中击败了65.14%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		nums := utils.StringToInts(scanner.Text())
		Printf("Output: %v\n", singleNumber(nums))
	}
}
