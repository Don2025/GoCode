package main

import (
	"bufio"
	. "fmt"
	"os"
	"strings"
)

// https://leetcode.cn/problems/longest-harmonious-subsequence/
//------------------------Leetcode Problem 594------------------------
func findLHS(nums []int) int {
	var ans int
	cnt := make(map[int]int)
	for _, num := range nums {
		cnt[num]++
	}
	for num, c := range cnt {
		if cnt[num+1] > 0 {
			ans = max(ans, c+cnt[num+1])
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

//------------------------Leetcode Problem 594------------------------
/*
 * https://leetcode.cn/problems/longest-harmonious-subsequence/
 * 执行用时：56ms 在所有Go提交中击败了60.98%的用户
 * 占用内存：7.2MB 在所有Go提交中击败了25.00%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		nums := strings.Fields(scanner.Text())
		Printf("Output: %d\n", findLHS(make([]int, len(nums))))
	}
}
