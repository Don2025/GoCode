package main

import (
	"bufio"
	. "fmt"
	"github.com/Don2025/GoCode/utils"
	"os"
	"sort"
)

// https://leetcode.cn/problems/3sum/
func threeSum(nums []int) [][]int {
	n := len(nums)
	sort.Ints(nums)
	var ans [][]int
	for i := 0; i < n; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		k := n - 1
		target := -1 * nums[i]
		for j := i + 1; j < n; j++ {
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}
			for j < k && nums[j]+nums[k] > target {
				k--
			}
			if j == k {
				break
			}
			if nums[j]+nums[k] == target {
				ans = append(ans, []int{nums[i], nums[j], nums[k]})
			}
		}
	}
	return ans
}

/*
 * https://leetcode.cn/problems/3sum/
 * 执行用时：32ms 在所有Go提交中击败了66.36%的用户
 * 占用内存：7.3MB 在所有Go提交中击败了68.63%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		nums := utils.StringToInts(scanner.Text())
		Printf("%v\n", threeSum(nums))
	}
}
