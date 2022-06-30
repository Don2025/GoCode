package main

import (
	"bufio"
	. "fmt"
	"github.com/Don2025/GoCode/utils"
	"os"
	"sort"
	"strconv"
)

// https://leetcode.cn/problems/non-overlapping-intervals/
//------------------------Leetcode Problem 435------------------------
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

//------------------------Leetcode Problem 435------------------------
/*
 * https://leetcode.cn/problems/non-overlapping-intervals/
 * 执行用时：204ms 在所有Go提交中击败了79.28%的用户
 * 占用内存：15.6MB 在所有Go提交中击败了90.295%的用户
**/

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		n, _ := strconv.Atoi(input.Text())
		intervals := make([][]int, n)
		for i := range intervals {
			input.Scan()
			intervals[i] = utils.StringToInts(input.Text())
		}
		Printf("Output: %v\n", eraseOverlapIntervals(intervals))
	}
}
