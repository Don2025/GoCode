package main

import (
	"bufio"
	. "fmt"
	"github.com/Don2025/GoCode/utils"
	"os"
	"sort"
)

// https://leetcode.cn/problems/sort-array-by-increasing-frequency/
// ------------------------Leetcode Problem 1636------------------------

func frequencySort(nums []int) []int {
	cnt := map[int]int{}
	for _, x := range nums {
		cnt[x]++
	}
	sort.Slice(nums, func(i, j int) bool {
		a, b := nums[i], nums[j]
		return cnt[a] < cnt[b] || cnt[a] == cnt[b] && a > b
	})
	return nums
}

// ------------------------Leetcode Problem 1636------------------------
/*
 * https://leetcode.cn/problems/sort-array-by-increasing-frequency/
 * 执行用时：4ms 在所有Go提交中击败了87.50%的用户
 * 占用内存：3MB 在所有Go提交中击败了70.00%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		nums := utils.StringToInts(scanner.Text())
		Printf("Output: %v\n", frequencySort(nums))
	}
}
