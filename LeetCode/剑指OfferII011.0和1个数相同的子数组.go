package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func findMaxLength(nums []int) int {
	m := map[int]int{0: -1}
	var cnt, ans int
	for i, x := range nums {
		if x == 1 {
			cnt++
		} else {
			cnt--
		}
		if pre, ok := m[cnt]; ok {
			ans = max(ans, i-pre)
		} else {
			m[cnt] = i
		}
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		nums := stringArrayToIntArray(strings.Fields(input.Text()))
		println(findMaxLength(nums))
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
 * 执行用时：84ms 在所有Go提交中击败了83.22%的用户
 * 占用内存：7.2MB 在所有Go提交中击败了49.42%的用户
**/
