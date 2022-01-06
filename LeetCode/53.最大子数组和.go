package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func maxSubArray(nums []int) int {
	ans := nums[0]
	var sum int
	for _, num := range nums {
		if sum < 0 {
			sum = 0
		}
		sum += num
		ans = max(ans, sum)
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		nums := stringArrayToIntArray(strings.Fields(input.Text()))
		println(maxSubArray(nums))
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
 * 执行用时：96ms 在所有Go提交中击败了81.81%的用户
 * 占用内存：9.3MB 在所有Go提交中击败了31.36%的用户
**/
