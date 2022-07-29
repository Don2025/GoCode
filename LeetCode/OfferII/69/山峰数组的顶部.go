package main

import (
	"bufio"
	. "fmt"
	"github.com/Don2025/GoCode/utils"
	"os"
)

// https://leetcode.cn/problems/B1IidL/
//-------------------------剑指 Offer II Problem 69------------------------
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

//-------------------------剑指 Offer II Problem 69------------------------
/*
 * https://leetcode.cn/problems/B1IidL/
 * 执行用时：8ms 在所有Go提交中击败了88.24%的用户
 * 占用内存：4.5MB 在所有Go提交中击败了100.00%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		arr := utils.StringToInts(scanner.Text())
		Printf("Output: %v\n", peakIndexInMountainArray(arr))
	}
}
