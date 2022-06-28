package main

import (
	"bufio"
	. "fmt"
	"github.com/Don2025/GoCode/utils"
	"math"
	"os"
)

// https://leetcode.cn/problems/increasing-triplet-subsequence/
//------------------------Leetcode Problem 334------------------------
func increasingTriplet(nums []int) bool {
	n := len(nums)
	if n < 3 {
		return false
	}
	first, second := nums[0], math.MaxInt32
	for i := 1; i < n; i++ {
		num := nums[i]
		if num > second {
			return true
		} else if num > first {
			second = num
		} else {
			first = num
		}
	}
	return false
}

/*
 * https://leetcode.cn/problems/increasing-triplet-subsequence/
 * 执行用时：56ms 在所有Go提交中击败了94.63%的用户
 * 占用内存：17MB 在所有Go提交中击败了97.99%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		nums := utils.StringToInts(scanner.Text())
		Printf("Output: %v\n", increasingTriplet(nums))
	}
}
