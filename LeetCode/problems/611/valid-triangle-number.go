package main

import (
	"bufio"
	. "fmt"
	"github.com/Don2025/GoCode/utils"
	"os"
	"sort"
)

// https://leetcode.cn/problems/valid-triangle-number/
//------------------------Leetcode Problem 611------------------------
func triangleNumber(nums []int) int {
	sort.Ints(nums)
	n := len(nums)
	var ans int
	for i := range nums {
		k := i
		for j := i + 1; j < n; j++ {
			for k+1 < n && nums[k+1] < nums[i]+nums[j] {
				k++
			}
			ans += max(0, k-j)
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

//------------------------Leetcode Problem 611------------------------
/*
 * https://leetcode.cn/problems/valid-triangle-number/
 * 执行用时：40ms 在所有Go提交中击败了56.51%的用户
 * 占用内存：3MB 在所有Go提交中击败了53.90%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		nums := utils.StringToInts(scanner.Text())
		Printf("Output: %v\n", triangleNumber(nums))
	}
}
