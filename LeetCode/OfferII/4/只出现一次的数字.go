package main

import (
	"bufio"
	. "fmt"
	"github.com/Don2025/GoCode/utils"
	"os"
)

// https://leetcode.cn/problems/WGki4K/
// ------------------------剑指 Offer II Problem 4------------------------
func singleNumber(nums []int) int {
	m := map[int]int{}
	for _, x := range nums {
		m[x]++
	}
	for k, v := range m {
		if v == 1 {
			return k
		}
	}
	return 0
}

// ------------------------剑指 Offer II Problem 4------------------------
/*
 * https://leetcode.cn/problems/WGki4K/
 * 执行用时：4ms 在所有Go提交中击败了92.99%的用户
 * 占用内存：4MB 在所有Go提交中击败了14.81%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		nums := utils.StringToInts(scanner.Text())
		Printf("Output: %v\n", singleNumber(nums))
	}
}
