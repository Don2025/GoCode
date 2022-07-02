package main

import (
	"bufio"
	. "fmt"
	"github.com/Don2025/GoCode/utils"
	"os"
	"strconv"
)

// https://leetcode.cn/problems/subarray-sum-equals-k/
//------------------------Leetcode Problem 560------------------------
func subarraySum(nums []int, k int) int {
	var ans int
	for i := 0; i < len(nums); i++ {
		var sum int
		for j := i; j >= 0; j-- {
			sum += nums[j]
			if sum == k {
				ans++
			}
		}
	}
	return ans
}

//------------------------Leetcode Problem 560------------------------
/*
 * https://leetcode.cn/problems/subarray-sum-equals-k/
 * 执行用时：1096ms 在所有Go提交中击败了6.60%的用户
 * 占用内存：6.5MB 在所有Go提交中击败了80.97%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		nums := utils.StringToInts(scanner.Text())
		scanner.Scan()
		k, _ := strconv.Atoi(scanner.Text())
		Printf("Output: %v\n", subarraySum(nums, k))
	}
}
