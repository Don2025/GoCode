package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func minSubArrayLen(target int, nums []int) int {
	var i, sum, ans int
	for j := 0; j < len(nums); j++ {
		sum += nums[j]
		for sum >= target {
			if ans == 0 {
				ans = j - i + 1
			} else {
				ans = min(ans, j-i+1)
			}
			sum -= nums[i]
			i++
		}
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
 * 执行用时：4ms 在所有Go提交中击败了96.58%的用户
 * 占用内存：3.6MB 在所有Go提交中击败了100.00%的用户
**/
