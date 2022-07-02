package main

import (
	"bufio"
	. "fmt"
	"github.com/Don2025/GoCode/utils"
	"os"
	"strconv"
)

// https://leetcode.cn/problems/continuous-subarray-sum/
//------------------------Leetcode Problem 523------------------------
func checkSubarraySum(nums []int, k int) bool {
	mp := map[int]int{0: -1}
	var mod int
	for i, x := range nums {
		mod = (mod + x) % k
		if preidx, has := mp[mod]; has {
			if i-preidx >= 2 {
				return true
			}
		} else {
			mp[mod] = i
		}
	}
	return false
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		nums := utils.StringToInts(scanner.Text())
		scanner.Scan()
		n, _ := strconv.Atoi(scanner.Text())
		Printf("Output: %v\n", checkSubarraySum(nums, n))
	}
}

//------------------------Leetcode Problem 523------------------------
/*
 * https://leetcode.cn/problems/continuous-subarray-sum/
 * 执行用时：228ms 在所有Go提交中击败了5.04%的用户
 * 占用内存：9.2MB 在所有Go提交中击败了38.13%的用户
**/
