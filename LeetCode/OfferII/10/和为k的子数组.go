package main

import (
	"bufio"
	. "fmt"
	"github.com/Don2025/GoCode/utils"
	"os"
	"strconv"
)

// https://leetcode.cn/problems/QTMn0o/
// ------------------------剑指 Offer II Problem 10------------------------
func subarraySum(nums []int, k int) int {
	m := make(map[int]int)
	m[0] = 1
	var pre, ans int
	for _, x := range nums {
		pre += x
		if _, ok := m[pre-k]; ok {
			ans += m[pre-k]
		}
		m[pre]++
	}
	return ans
}

// ------------------------剑指 Offer II Problem 10------------------------
/*
 * https://leetcode.cn/problems/QTMn0o/
 * 执行用时：36ms 在所有Go提交中击败了78.36%的用户
 * 占用内存：8.1MB 在所有Go提交中击败了5.04%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		nums := utils.StringToInts(scanner.Text())
		scanner.Scan()
		k, _ := strconv.Atoi(scanner.Text())
		Printf("Output: %v\n", subarraySum(nums, k))
	}
}
