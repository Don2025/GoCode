package main

import (
	"bufio"
	"os"
	"sort"
	"strconv"
	"strings"
)

func canMakeArithmeticProgression(arr []int) bool {
	if len(arr) <= 2 {
		return true
	}
	sort.Ints(arr)
	d := arr[1] - arr[0]
	for i := 1; i < len(arr)-1; i++ {
		if arr[i+1]-arr[i] != d {
			return false
		}
	}
	return true
}

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		arr := stringArrayToIntArray(strings.Fields(input.Text()))
		println(canMakeArithmeticProgression(arr))
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
 * 占用内存：2.4MB 在所有Go提交中击败了95.94%的用户
**/
