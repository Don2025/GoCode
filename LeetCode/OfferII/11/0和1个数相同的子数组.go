package main

import (
	"bufio"
	. "fmt"
	"github.com/Don2025/GoCode/utils"
	"os"
)

// https://leetcode.cn/problems/A1NYOS/
// ------------------------剑指 Offer II Problem 11------------------------
func findMaxLength(nums []int) int {
	m := map[int]int{0: -1}
	var cnt, ans int
	for i, x := range nums {
		if x == 1 {
			cnt++
		} else {
			cnt--
		}
		if pre, ok := m[cnt]; ok {
			ans = max(ans, i-pre)
		} else {
			m[cnt] = i
		}
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// ------------------------剑指 Offer II Problem 11------------------------
/*
 * https://leetcode.cn/problems/A1NYOS/
 * 执行用时：84ms 在所有Go提交中击败了83.22%的用户
 * 占用内存：7.2MB 在所有Go提交中击败了49.42%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		nums := utils.StringToInts(scanner.Text())
		Printf("Output: %v\n", findMaxLength(nums))
	}
}
