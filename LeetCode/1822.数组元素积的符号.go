package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func arraySign(nums []int) int {
	ans := 1
	for _, x := range nums {
		if x < 0 {
			ans = -ans
		} else if x == 0 {
			return 0
		}
	}
	return ans
}

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		nums := stringArrayToIntArray(strings.Fields(input.Text()))
		println(arraySign(nums))
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
 * 执行用时：4ms 在所有Go提交中击败了95.28%的用户
 * 占用内存：3.1MB 在所有Go提交中击败了73.23%的用户
**/
