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
	for _, x := range nums {
		if sum < 0 {
			sum = x
		} else {
			sum += x
		}
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
		println(maxSubArray(stringArrayToIntArray(strings.Fields(input.Text()))))
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
 * 执行用时：32ms 在所有Go提交中击败了13.43%的用户
 * 占用内存：7.1MB 在所有Go提交中击败了90.56%的用户
**/
