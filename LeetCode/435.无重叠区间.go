package main

import (
	"bufio"
	"os"
	"sort"
	"strconv"
	"strings"
)

func eraseOverlapIntervals(intervals [][]int) int {
	n := len(intervals)
	if n == 0 {
		return 0
	}
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][1] < intervals[j][1]
	})
	ans, right := 1, intervals[0][1]
	for _, p := range intervals[1:] {
		if p[0] >= right {
			ans++
			right = p[1]
		}
	}
	return n - ans
}

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		n, _ := strconv.Atoi(input.Text())
		intervals := make([][]int, n)
		for i := range intervals {
			input.Scan()
			intervals[i] = stringArrayToIntArray(strings.Fields(input.Text()))
		}
		println(eraseOverlapIntervals(intervals))
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
 * 执行用时：204ms 在所有Go提交中击败了79.28%的用户
 * 占用内存：15.6MB 在所有Go提交中击败了90.295%的用户
**/
