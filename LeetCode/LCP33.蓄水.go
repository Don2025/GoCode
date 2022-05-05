package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func storeWater(bucket []int, vat []int) int {
	m, n := 0, len(bucket)
	for i := 0; i < n; i++ {
		m = max(m, vat[i])
	}
	if m == 0 {
		return 0
	}
	ans := m + 1
	for i := 1; i < m; i++ {
		cnt := i
		for j := 0; j < n; j++ {
			cnt += max(0, (vat[j]+i-1)/i-bucket[j])
		}
		ans = min(ans, cnt)
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		bucket := stringArrayToIntArray(strings.Fields(input.Text()))
		input.Scan()
		vat := stringArrayToIntArray(strings.Fields(input.Text()))
		println(storeWater(bucket, vat))
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
 * 执行用时：60ms 在所有Go提交中击败了45.16%的用户
 * 占用内存：2.1MB 在所有Go提交中击败了77.42%的用户
**/
