package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func findPeakElement(nums []int) int {
	l, r := 0, len(nums)-1
	for l < r {
		m := l + (r-l)>>1
		if nums[m] > nums[m+1] {
			r = m
		} else {
			l = m + 1
		}
	}
	return l
}

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		println(findPeakElement(stringArrayToIntArray(strings.Fields(input.Text()))))
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
 * 执行用时：4ms 在所有Go提交中击败了30.27%的用户
 * 占用内存：2.7MB 在所有Go提交中击败了61.98%的用户
**/
