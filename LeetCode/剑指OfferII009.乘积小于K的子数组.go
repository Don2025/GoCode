package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func numSubarrayProductLessThanK(nums []int, k int) int {
	if k <= 1 {
		return 0
	}
	var ans int
	mul := 1
	for l, r := 0, 0; r < len(nums); r++ {
		mul *= nums[r]
		for mul >= k {
			mul /= nums[l]
			l++
		}
		ans += r - l + 1
	}
	return ans
}

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		nums := stringArrayToIntArray(strings.Fields(input.Text()))
		input.Scan()
		k, _ := strconv.Atoi(input.Text())
		println(numSubarrayProductLessThanK(nums, k))
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
 * 执行用时：64ms 在所有Go提交中击败了95.08%的用户
 * 占用内存：7.2MB 在所有Go提交中击败了66.67%的用户
**/
