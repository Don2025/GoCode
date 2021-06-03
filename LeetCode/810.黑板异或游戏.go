package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func xorGame(nums []int) bool {
	var flag, n = 0, len(nums)
	for _, x := range nums {
		flag ^= x
	}
	if flag == 0 || n%2 == 0 {
		return true
	}
	return false
}

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		a := stringArrayToIntArray(strings.Fields(input.Text()))
		println(xorGame(a))
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
 * 执行用时：8ms 在所有Go提交中击败了66.67%的用户
 * 占用内存：3.9MB 在所有Go提交中击败了100.00%的用户
**/
