package main

import (
	"bufio"
	. "fmt"
	"github.com/Don2025/GoCode/utils"
	"os"
)

// https://leetcode.cn/problems/WhsWhI/
// ------------------------剑指 Offer II Problem 119------------------------
func longestConsecutive(nums []int) int {
	numSet := map[int]bool{}
	for _, num := range nums {
		numSet[num] = true
	}
	longestStreak := 0
	for num := range numSet {
		if !numSet[num-1] {
			currentNum := num
			currentStreak := 1
			for numSet[currentNum+1] {
				currentNum++
				currentStreak++
			}
			if longestStreak < currentStreak {
				longestStreak = currentStreak
			}
		}
	}
	return longestStreak
}

// ------------------------剑指 Offer II Problem 119------------------------
/*
 * https://leetcode.cn/problems/WhsWhI/
 * 执行用时：32ms 在所有Go提交中击败了93.23%的用户
 * 占用内存：8.1MB 在所有Go提交中击败了62.78%的用户
**/
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	Printf("Input a line of numbers separated by space:")
	for scanner.Scan() {
		nums := utils.StringToInts(scanner.Text())
		Printf("Output: %v\n", longestConsecutive(nums))
		Printf("Input a line of numbers separated by space:")
	}
}
