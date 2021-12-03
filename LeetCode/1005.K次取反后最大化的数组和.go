package main

import (
	"bufio"
	"os"
	"sort"
	"strconv"
	"strings"
)

func largestSumAfterKNegations(nums []int, k int) int {
	sort.Slice(nums, func(i, j int) bool {
		return abs(nums[i]) > abs(nums[j])
	})
	for i := 0; i < len(nums); i++ {
		if k > 0 && nums[i] < 0 {
			nums[i] = abs(nums[i])
			k--
		}
	}
	if k&1 != 0 {
		nums[len(nums)-1] = -nums[len(nums)-1]
	}
	var ans int
	for i := 0; i < len(nums); i++ {
		ans += nums[i]
	}
	return ans
}

func abs(n int) int {
	if n > 0 {
		return n
	}
	return -n
}

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		nums := stringArrayToIntArray(strings.Fields(input.Text()))
		input.Scan()
		k, _ := strconv.Atoi(input.Text())
		println(largestSumAfterKNegations(nums, k))
	}
}

func stringArrayToIntArray(strings []string) []int {
	var arr []int
	for _, x := range strings {
		n, _ := strconv.Atoi(x)
		arr = append(arr, n)
	}
	return arr
}

/*
 * 执行用时：4ms 在所有Go提交中击败了94.59%的用户
 * 占用内存：2.6MB 在所有Go提交中击败了68.24%的用户
**/
