package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func minMoves(nums []int) int {
	var ans int
	min := nums[0]
	for _, num := range nums[1:] {
		if num < min {
			min = num
		}
	}
	for _, num := range nums {
		ans += num - min
	}
	return ans
}

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		println(minMoves(stringArrayToIntArray(strings.Fields(input.Text()))))
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
 * 执行用时：36ms 在所有Go提交中击败了66.06%的用户
 * 占用内存：6.6MB 在所有Go提交中击败了98.17%的用户
**/
