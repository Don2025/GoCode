package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func compareVersion(version1 string, version2 string) int {
	arr1, arr2 := strings.Split(version1, "."), strings.Split(version2, ".")
	for i := 0; i < max(len(arr1), len(arr2)); i++ {
		n1, n2 := getValue(arr1, i), getValue(arr2, i)
		if n1 < n2 {
			return -1
		} else if n1 > n2 {
			return 1
		}
	}
	return 0
}

func getValue(arr []string, n int) int {
	if n < len(arr) {
		ans, _ := strconv.Atoi(arr[n])
		return ans
	}
	return 0
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
		version1 := input.Text()
		input.Scan()
		version2 := input.Text()
		println(compareVersion(version1, version2))
	}
}

/*
 * 执行用时：0ms 在所有Go提交中击败了100.00%的用户
 * 占用内存：1.9MB 在所有Go提交中击败了79.15%的用户
**/
