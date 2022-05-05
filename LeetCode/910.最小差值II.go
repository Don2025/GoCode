package main

import (
	"bufio"
	"os"
	"sort"
	"strconv"
	"strings"
)

func smallestRangeII(nums []int, k int) int {
	sort.Ints(nums)
	n := len(nums)
	ans := nums[n-1] - nums[0]
	for i := 0; i < len(nums)-1; i++ {
		high := max(nums[n-1]-k, nums[i]+k)
		low := min(nums[0]+k, nums[i+1]-k)
		ans = min(ans, high-low)
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		nums := stringArrayToIntArray(strings.Fields(input.Text()))
		input.Scan()
		k, _ := strconv.Atoi(input.Text())
		println(smallestRangeII(nums, k))
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
 * 执行用时：20ms 在所有Go提交中击败了77.97%的用户
 * 占用内存：5.9MB 在所有Go提交中击败了83.05%的用户
**/
