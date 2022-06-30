package main

import (
	"bufio"
	. "fmt"
	"github.com/Don2025/GoCode/utils"
	"os"
	"sort"
)

// https://leetcode.cn/problems/minimum-moves-to-equal-array-elements-ii/
//------------------------Leetcode Problem 462------------------------
func minMoves2(nums []int) int {
	sort.Ints(nums)
	n := nums[len(nums)>>1]
	var ans int
	for _, x := range nums {
		ans += abs(x - n)
	}
	return ans
}

func abs(n int) int {
	if n < 0 {
		n = -n
	}
	return n
}

//------------------------Leetcode Problem 462------------------------
/*
 * https://leetcode.cn/problems/minimum-moves-to-equal-array-elements-ii/
 * 执行用时：8ms 在所有Go提交中击败了86.71%的用户
 * 占用内存：4MB 在所有Go提交中击败了100.00%的用户
**/

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		nums := utils.StringToInts(input.Text())
		Printf("Output: %v\n", minMoves2(nums))
	}
}
