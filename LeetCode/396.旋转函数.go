package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func maxRotateFunction(nums []int) int {
	var f, sum int
	for i := range nums {
		sum += nums[i]
		f += i * nums[i]
	}
	ans, n := f, len(nums)
	for i := 1; i < n; i++ {
		f += sum - n*nums[n-i]
		ans = max(ans, f)
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
		println(maxRotateFunction(nums))
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
 * 执行用时：120ms 在所有Go提交中击败了91.43%的用户
 * 占用内存：7.9MB 在所有Go提交中击败了65.71%的用户
**/
