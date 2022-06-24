package main

import (
	"bufio"
	. "fmt"
	"github.com/Don2025/GoCode/utils"
	"os"
)

// https://leetcode.cn/problems/subsets/
//------------------------Leetcode Problem 78------------------------
func subsets(nums []int) [][]int {
	n := len(nums)
	var ans [][]int
	for i := 0; i < 1<<n; i++ {
		var arr []int
		for j := range nums {
			if i>>j&1 > 0 {
				arr = append(arr, nums[j])
			}
		}
		ans = append(ans, append([]int(nil), arr...))
	}
	return ans
}

//------------------------Leetcode Problem 78------------------------
/*
 * https://leetcode.cn/problems/subsets/
 * 执行用时：0ms 在所有Go提交中击败了100.00%的用户
 * 占用内存：2.2MB 在所有Go提交中击败了25.72%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		nums := utils.StringToInts(scanner.Text())
		Printf("Output: %v\n", subsets(nums))
	}
}
