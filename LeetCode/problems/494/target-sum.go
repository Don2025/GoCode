package main

import (
	"bufio"
	. "fmt"
	"github.com/Don2025/GoCode/utils"
	"os"
	"strconv"
)

// https://leetcode.cn/problems/target-sum/
//------------------------Leetcode Problem 494------------------------
func findTargetSumWays(nums []int, target int) int {
	sum, n := 0, len(nums)
	for _, x := range nums {
		sum += x
	}
	t := sum - target
	if t%2 != 0 || t < 0 {
		return 0
	}
	t /= 2
	f := make([]int, t+1)
	f[0] = 1
	for i := 1; i <= n; i++ {
		for j := t; j >= nums[i-1]; j-- {
			f[j] += f[j-nums[i-1]]
		}
	}
	return f[t]
}

//------------------------Leetcode Problem 494------------------------
/*
 * https://leetcode.cn/problems/target-sum/
 * 执行用时：0ms 在所有Go提交中击败了100.00%的用户
 * 占用内存：2.4MB 在所有Go提交中击败了60.25%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		nums := utils.StringToInts(scanner.Text())
		scanner.Scan()
		target, _ := strconv.Atoi(scanner.Text())
		Printf("Output: %v\n", findTargetSumWays(nums, target))
	}
}
