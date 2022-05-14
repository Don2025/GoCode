package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func subarraySum(nums []int, k int) int {
	m := make(map[int]int)
	m[0] = 1
	var pre, ans int
	for _, x := range nums {
		pre += x
		if _, ok := m[pre-k]; ok {
			ans += m[pre-k]
		}
		m[pre]++
	}
	return ans
}

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		nums := stringArrayToIntArray(strings.Fields(input.Text()))
		input.Scan()
		k, _ := strconv.Atoi(input.Text())
		println(subarraySum(nums, k))
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
 * 执行用时：36ms 在所有Go提交中击败了78.36%的用户
 * 占用内存：8.1MB 在所有Go提交中击败了5.04%的用户
**/
