package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func sumOddLengthSubarrays(arr []int) int {
	n := len(arr)
	prefix := make([]int, n+1)
	for i := 0; i < n; i++ {
		prefix[i+1] = prefix[i] + arr[i]
	}
	var ans int
	for i := 0; i < n; i++ {
		for j := i; j >= 0; j -= 2 {
			ans += prefix[i+1] - prefix[j]
		}
	}
	return ans
}

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		arr := stringArrayToIntArray(strings.Fields(input.Text()))
		println(sumOddLengthSubarrays(arr))
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
 * 执行用时：0ms 在所有Go提交中击败了100.00%的用户
 * 占用内存：2.1MB 在所有Go提交中击败了43.59%的用户
**/
