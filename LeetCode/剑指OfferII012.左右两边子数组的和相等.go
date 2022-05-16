package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func pivotIndex(nums []int) int {
	var sum, cur int
	for _, x := range nums {
		sum += x
	}
	for i := range nums {
		sum -= nums[i]
		if cur == sum {
			return i
		}
		cur += nums[i]
	}
	return -1
}

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		nums := stringArrayToIntArray(strings.Fields(input.Text()))
		println(pivotIndex(nums))
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
 * 执行用时：24ms 在所有Go提交中击败了11.55%的用户
 * 占用内存：6MB 在所有Go提交中击败了72.52%的用户
**/
