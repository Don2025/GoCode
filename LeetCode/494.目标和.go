package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func findTargetSumWays(nums []int, target int) int {
	sum, n := 0, len(nums)
	for _, x := range nums {
		sum += x
	}
	t := sum - target
	if t%2 != 0 || t < 0 {
		return 0
	}
	t /= 2
	f := make([]int, t+1)
	f[0] = 1
	for i := 1; i <= n; i++ {
		for j := t; j >= nums[i-1]; j-- {
			f[j] += f[j-nums[i-1]]
		}
	}
	return f[t]
}

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		nums := stringArrayToIntArray(strings.Fields(input.Text()))
		input.Scan()
		target, _ := strconv.Atoi(input.Text())
		println(findTargetSumWays(nums, target))
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
 * 执行用时：0ms 在所有Go提交中击败了100.00%的用户
 * 占用内存：2.4MB 在所有Go提交中击败了60.25%的用户
**/
