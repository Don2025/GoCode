package main

import (
	"bufio"
	"os"
	"sort"
	"strconv"
	"strings"
)

func hIndex(citations []int) int {
	sort.Ints(citations)
	for i := 0; i < len(citations); i++ {
		h := len(citations) - i
		if h <= citations[i] {
			return h
		}
	}
	return 0
}

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		println(hIndex(stringArrayToIntArray(strings.Fields(input.Text()))))
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
 * 占用内存：2.2MB 在所有Go提交中击败了97.32%的用户
**/
