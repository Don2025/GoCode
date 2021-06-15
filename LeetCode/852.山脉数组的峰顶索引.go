package main

import (
	"bufio"
	"os"
	"sort"
	"strconv"
	"strings"
)

func peakIndexInMountainArray(arr []int) int {
	return sort.Search(len(arr)-1, func(i int) bool {
		return arr[i] > arr[i+1]
	})
}

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		println(peakIndexInMountainArray(stringArrayToIntArray(strings.Fields(input.Text()))))
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
 * 执行用时：16ms 在所有Go提交中击败了7.24%的用户
 * 占用内存：4.5MB 在所有Go提交中击败了63.82%的用户
**/
