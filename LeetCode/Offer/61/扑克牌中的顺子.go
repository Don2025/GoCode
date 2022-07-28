package main

import (
	"bufio"
	. "fmt"
	"github.com/Don2025/GoCode/utils"
	"os"
	"sort"
)

// https://leetcode.cn/problems/bu-ke-pai-zhong-de-shun-zi-lcof/
// ------------------------剑指 Offer I Problem 61------------------------
func isStraight(nums []int) bool {
	var joker int
	sort.Ints(nums)
	for i := 0; i < 4; i++ {
		if nums[i] == 0 {
			joker++ //大小王
		} else if nums[i] == nums[i+1] {
			return false
		}
	}
	return nums[4]-nums[joker] < 5
}

// ------------------------剑指 Offer I Problem 61------------------------
/*
 * https://leetcode.cn/problems/bu-ke-pai-zhong-de-shun-zi-lcof/
 * 执行用时：0ms 在所有Go提交中击败了100.00%的用户
 * 占用内存：1.8MB 在所有Go提交中击败了75.21%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		nums := utils.StringToInts(scanner.Text())
		Printf("Output: %v\n", isStraight(nums))
	}
}
