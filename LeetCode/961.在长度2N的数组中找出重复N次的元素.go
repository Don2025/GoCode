package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func repeatedNTimes(nums []int) int {
	m := make(map[int]bool)
	for _, x := range nums {
		if m[x] {
			return x
		}
		m[x] = true
	}
	return -1
}

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		nums := stringArrayToIntArray(strings.Fields(input.Text()))
		println(repeatedNTimes(nums))
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
 * 执行用时：20ms 在所有Go提交中击败了96.60%的用户
 * 占用内存：6.5MB 在所有Go提交中击败了75.90%的用户
**/
