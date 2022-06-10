package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

/* 在一个长度为 n 的数组 nums 里的所有数字都在 0～n-1 的范围内。
 * 数组中某些数字是重复的，但不知道有几个数字重复了，也不知道每个数字重复了几次。请找出数组中任意一个重复的数字。
**/
func findRepeatNumber(nums []int) int {
	a := make([]bool, len(nums))
	for _, x := range nums {
		if !a[x] {
			a[x] = true
		} else {
			return x
		}
	}
	return -1
}

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		nums := stringArrayToIntArray(strings.Fields(input.Text()))
		println(findRepeatNumber(nums))
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
 * 执行用时：40ms 在所有Go提交中击败了89.12%的用户
 * 占用内存：7.9MB 在所有Go提交中击败了93.47%的用户
**/
