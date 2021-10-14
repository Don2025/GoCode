package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func peakIndexInMountainArray(arr []int) int {
	l, r := 1, len(arr)-1
	for l <= r {
		m := l + (r-l)>>1
		if arr[m] > arr[m-1] {
			if arr[m] > arr[m+1] {
				return m
			} else {
				l = m + 1
			}
		} else {
			r = m - 1
		}
	}
	return 0
}

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		println(peakIndexInMountainArray(stringArrayToIntArray(strings.Fields(input.Text()))))
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
 * 执行用时：8ms 在所有Go提交中击败了88.24%的用户
 * 占用内存：4.5MB 在所有Go提交中击败了100.00%的用户
**/
