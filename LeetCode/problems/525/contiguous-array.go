package main

import (
	"bufio"
	. "fmt"
	"github.com/Don2025/GoCode/utils"
	"os"
)

// https://leetcode.cn/problems/contiguous-array/
//------------------------Leetcode Problem 525------------------------
func findMaxLength(nums []int) int {
	var cnt, ans int
	mp := map[int]int{0: -1}
	for i := 0; i < len(nums); i++ {
		if nums[i] == 1 {
			cnt++
		} else {
			cnt--
		}
		if preidx, has := mp[cnt]; has {
			t := i - preidx
			if t > ans {
				ans = t
			}
		} else {
			mp[cnt] = i
		}
	}
	return ans
}

//------------------------Leetcode Problem 525------------------------
/*
 * https://leetcode.cn/problems/contiguous-array/
 * 执行用时：132ms 在所有Go提交中击败了22.22%的用户
 * 占用内存：7.5MB 在所有Go提交中击败了27.78%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		nums := utils.StringToInts(scanner.Text())
		Printf("Output: %v\n", findMaxLength(nums))
	}
}
