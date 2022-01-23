package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func getMaxLen(nums []int) int {
	var pos, neg int
	if nums[0] > 0 {
		pos = 1
	}
	if nums[0] < 0 {
		neg = 1
	}
	ans := pos
	for i := 1; i < len(nums); i++ {
		if nums[i] > 0 {
			pos++
			if neg > 0 {
				neg++
			} else {
				neg = 0
			}
		} else if nums[i] < 0 {
			t := pos
			if neg > 0 {
				pos = neg + 1
			} else {
				pos = 0
			}
			neg = t + 1
		} else {
			pos, neg = 0, 0
		}
		ans = max(ans, pos)
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
		println(getMaxLen(nums))
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
 * 执行用时：100ms 在所有Go提交中击败了22.06%的用户
 * 占用内存：9.9MB 在所有Go提交中击败了84.31%的用户
**/
