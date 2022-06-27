package main

import (
	"bufio"
	. "fmt"
	"github.com/Don2025/GoCode/utils"
	"os"
	"strconv"
)

// https://leetcode.cn/problems/two-sum-ii-input-array-is-sorted/
//------------------------Leetcode Problem 167------------------------
func twoSum(numbers []int, target int) []int {
	l, r := 0, len(numbers)-1
	for l < r {
		sum := numbers[l] + numbers[r]
		if target == sum {
			return []int{l + 1, r + 1}
		} else if target < sum {
			r--
		} else { // target > sum
			l++
		}
	}
	return []int{}
}

//------------------------Leetcode Problem 167------------------------
/*
 * https://leetcode.cn/problems/two-sum-ii-input-array-is-sorted/
 * 执行用时：4ms 在所有Go提交中击败了86.97%的用户
 * 占用内存：3MB 在所有Go提交中击败了67.99%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		nums := utils.StringToInts(scanner.Text())
		scanner.Scan()
		target, _ := strconv.Atoi(scanner.Text())
		Printf("Output: %v\n", twoSum(nums, target))
	}
}
