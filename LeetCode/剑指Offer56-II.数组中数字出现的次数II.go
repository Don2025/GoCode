package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func singleNumber(nums []int) int {
	var a, b int
	for _, c := range nums {
		a, b = a&^c|b&c, b&^c|^a&^b&c
	}
	return ^a & b
}

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		nums := stringArrayToIntArray(strings.Fields(input.Text()))
		println(singleNumber(nums))
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
 * 执行用时：16ms 在所有Go提交中击败了99.24%的用户
 * 占用内存：6.3MB 在所有Go提交中击败了65.14%的用户
**/
