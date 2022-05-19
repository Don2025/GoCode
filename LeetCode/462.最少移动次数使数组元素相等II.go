package main

import (
	"bufio"
	"os"
	"sort"
	"strconv"
	"strings"
)

func minMoves2(nums []int) int {
	sort.Ints(nums)
	n := nums[len(nums)>>1]
	var ans int
	for _, x := range nums {
		ans += abs(x - n)
	}
	return ans
}

func abs(n int) int {
	if n < 0 {
		n = -n
	}
	return n
}

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		nums := stringArrayToIntArray(strings.Fields(input.Text()))
		println(minMoves2(nums))
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
 * 执行用时：8ms 在所有Go提交中击败了86.71%的用户
 * 占用内存：4MB 在所有Go提交中击败了100.00%的用户
**/
