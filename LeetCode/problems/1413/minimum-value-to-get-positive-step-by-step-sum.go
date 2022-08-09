package main

import (
	"bufio"
	. "fmt"
	"github.com/Don2025/GoCode/utils"
	"os"
)

// https://leetcode.cn/problems/minimum-value-to-get-positive-step-by-step-sum/
//------------------------Leetcode Problem 1413------------------------
func minStartValue(nums []int) int {
	accSum, accSumMin := 0, 0
	for _, num := range nums {
		accSum += num
		accSumMin = min(accSumMin, accSum)
	}
	return -accSumMin + 1
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

//------------------------Leetcode Problem 1413------------------------
/*
 * https://leetcode.cn/problems/minimum-value-to-get-positive-step-by-step-sum/
 * 执行用时：0ms 在所有Go提交中击败了100.00%的用户
 * 占用内存：1.9MB 在所有Go提交中击败了100.00%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		nums := utils.StringToInts(scanner.Text())
		Printf("Output: %v\n", minStartValue(nums))
	}
}
