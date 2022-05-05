package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func isMonotonic(nums []int) bool {
	n := len(nums)
	if n == 1 {
		return true
	}
	var a, b int
	for i := 0; i < n-1; i++ {
		if nums[i] < nums[i+1] {
			a = 1
		} else if nums[i] > nums[i+1] {
			b = 1
		}
	}
	if a+b == 2 {
		return false
	}
	return true
}

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		nums := stringArrayToIntArray(strings.Fields(input.Text()))
		println(isMonotonic(nums))
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
 * 执行用时：124ms 在所有Go提交中击败了57.87%的用户
 * 占用内存：8.7MB 在所有Go提交中击败了88.76%的用户
**/
