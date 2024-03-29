package main

import (
	"bufio"
	. "fmt"
	"os"
	"sort"
	"strconv"
)

// https://leetcode.cn/problems/reordered-power-of-2/
//------------------------Leetcode Problem 869------------------------
func reorderedPowerOf2(n int) bool {
	nums := []byte(strconv.Itoa(n))
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	vis := make([]bool, len(nums))
	var backtrack func(int, int) bool
	backtrack = func(index, num int) bool {
		if index == len(nums) {
			return isPowerOfTwo(num)
		}
		for i, x := range nums {
			if num == 0 && x == '0' || vis[i] || i > 0 && !vis[i-1] && x == nums[i-1] {
				continue
			}
			vis[i] = true
			if backtrack(index+1, num*10+int(x-'0')) {
				return true
			}
			vis[i] = false
		}
		return false
	}
	return backtrack(0, 0)
}

func isPowerOfTwo(n int) bool {
	return n&(n-1) == 0
}

//------------------------Leetcode Problem 869------------------------
/*
 * https://leetcode.cn/problems/reordered-power-of-2/
 * 执行用时：24ms 在所有Go提交中击败了47.37%的用户
 * 占用内存：2MB 在所有Go提交中击败了57.89%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		n, _ := strconv.Atoi(scanner.Text())
		Printf("Output: %t\n", reorderedPowerOf2(n))
	}
}
