package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func singleNumber(nums []int) int {
	var ans int
	for _, x := range nums {
		ans ^= x
	}
	return ans
}

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		println(singleNumber(stringArrayToIntArray(strings.Fields(input.Text()))))
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
 * 执行用时：12ms 在所有Go提交中击败了93.66%的用户
 * 占用内存：6MB 在所有Go提交中击败了78.66%的用户
**/
