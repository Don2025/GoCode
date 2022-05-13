package main

import (
	"bufio"
	"math"
	"os"
	"strconv"
	"strings"
)

func minSubArrayLen(target int, nums []int) int {
	ans := math.MaxInt32
	var sum int
	for i, j := 0, 0; j < len(nums); j++ {
		sum += nums[j]
		for i <= j && sum >= target {
			ans = min(ans, j-i+1)
			sum -= nums[i]
			i++
		}
	}
	if ans == math.MaxInt32 {
		ans = 0
	}
	return ans
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
		target, _ := strconv.Atoi(input.Text())
		input.Scan()
		nums := stringArrayToIntArray(strings.Fields(input.Text()))
		println(minSubArrayLen(target, nums))
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
 * 执行用时：4ms 在所有Go提交中击败了95.41%的用户
 * 占用内存：3.6MB 在所有Go提交中击败了70.64%的用户
**/
