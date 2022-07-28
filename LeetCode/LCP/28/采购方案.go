package main

import (
	"bufio"
	. "fmt"
	"github.com/Don2025/GoCode/utils"
	"os"
	"sort"
	"strconv"
)

// https://leetcode.cn/problems/4xy4Wx/
// ------------------------LeetCode Cup Problem 28------------------------
func purchasePlans(nums []int, target int) int {
	const mod = 1e9 + 7
	sort.Ints(nums)
	l, r := 0, len(nums)-1
	var sum int
	for l != r {
		if nums[l]+nums[r] > target {
			r--
		} else {
			sum += r - l
			l++
		}
	}
	return sum % mod
}

// ------------------------LeetCode Cup Problem 28------------------------
/*
 * https://leetcode.cn/problems/4xy4Wx/
 * 执行用时：136ms 在所有Go提交中击败了77.27%的用户
 * 占用内存：7.8MB 在所有Go提交中击败了78.79%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		nums := utils.StringToInts(scanner.Text())
		scanner.Scan()
		target, _ := strconv.Atoi(scanner.Text())
		Printf("Output: %d\n", purchasePlans(nums, target))
	}
}
