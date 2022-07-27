package main

import (
	"bufio"
	. "fmt"
	"github.com/Don2025/GoCode/utils"
	"os"
)

// https://leetcode.cn/problems/count-special-quadruplets/
//------------------------Leetcode Problem 1995------------------------
func countQuadruplets(nums []int) int {
	cnt := make(map[int]int)
	var ans int
	for i := len(nums) - 3; i > 0; i-- {
		for _, x := range nums[i+2:] {
			cnt[x-nums[i+1]]++
		}
		for _, x := range nums[:i] {
			if sum := x + nums[i]; cnt[sum] > 0 {
				ans += cnt[sum]
			}
		}
	}
	return ans
}

//------------------------Leetcode Problem 1995------------------------
/*
 * https://leetcode.cn/problems/count-special-quadruplets/
 * 执行用时：0ms 在所有Go提交中击败了100.00%的用户
 * 占用内存：2MB 在所有Go提交中击败了56.24%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		nums := utils.StringToInts(scanner.Text())
		Printf("Output: %d\n", countQuadruplets(nums))
	}
}
