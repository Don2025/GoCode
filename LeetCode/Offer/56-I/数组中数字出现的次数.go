package main

import (
	"bufio"
	. "fmt"
	"github.com/Don2025/GoCode/utils"
	"os"
)

func singleNumbers(nums []int) []int {
	m := make(map[int]int, len(nums))
	for _, x := range nums {
		if _, has := m[x]; !has {
			m[x] = x
		} else {
			delete(m, x)
		}
	}
	var ans []int
	for x := range m {
		ans = append(ans, x)
	}
	return ans
}

// ------------------------剑指 Offer I Problem 56-I------------------------
/*
 * https://leetcode.cn/problems/shu-zu-zhong-shu-zi-chu-xian-de-ci-shu-lcof/
 * 执行用时：12ms 在所有Go提交中击败了97.34%的用户
 * 占用内存：6.6MB 在所有Go提交中击败了65.47%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		nums := utils.StringToInts(scanner.Text())
		Printf("Output: %v\n", singleNumbers(nums))
	}
}
