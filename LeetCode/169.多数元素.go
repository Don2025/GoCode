package main

import (
	"bufio"
	"os"
	"sort"
	"strconv"
	"strings"
)

func majorityElement(nums []int) int {
	sort.Ints(nums)
	return nums[len(nums)>>1]
}

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		nums := stringArrayToIntArray(strings.Fields(input.Text()))
		println(majorityElement(nums))
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
 * 执行用时：12ms 在所有Go提交中击败了94.07%的用户
 * 占用内存：5.9MB 在所有Go提交中击败了51.98%的用户
**/
