package main

import (
	"bufio"
	. "fmt"
	"github.com/Don2025/GoCode/utils"
	"os"
	"strconv"
)

// https://leetcode.cn/problems/kLl5u1/
// ------------------------剑指 Offer II Problem 6------------------------
func twoSum(numbers []int, target int) []int {
	i, j := 0, len(numbers)-1
	for i < j {
		if numbers[i]+numbers[j] == target {
			break
		} else if numbers[i]+numbers[j] < target {
			i++
		} else {
			j--
		}
	}
	return []int{i, j}
}

// ------------------------剑指 Offer II Problem 6------------------------
/*
 * https://leetcode.cn/problems/kLl5u1/
 * 执行用时：4ms 在所有Go提交中击败了67.33%的用户
 * 占用内存：2.8MB 在所有Go提交中击败了100.00%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		numbers := utils.StringToInts(scanner.Text())
		scanner.Scan()
		target, _ := strconv.Atoi(scanner.Text())
		Printf("Output: %v\n", twoSum(numbers, target))
	}
}
