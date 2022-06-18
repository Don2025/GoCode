package main

import (
	"bufio"
	. "fmt"
	"github.com/Don2025/GoCode/utils"
	"os"
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
		arr := utils.StringToInts(input.Text())
		Printf("Output: %v\n", peakIndexInMountainArray(arr))
	}
}

/*
 * 执行用时：8ms 在所有Go提交中击败了88.24%的用户
 * 占用内存：4.5MB 在所有Go提交中击败了100.00%的用户
**/
