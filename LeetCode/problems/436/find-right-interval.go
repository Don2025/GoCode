package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

// https://leetcode.cn/problems/find-right-interval/
//------------------------Leetcode Problem 436------------------------
func findRightInterval(intervals [][]int) []int {
	for i := range intervals {
		intervals[i] = append(intervals[i], i)
	}
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	n := len(intervals)
	ans := make([]int, n)
	for _, p := range intervals {
		i := sort.Search(n, func(i int) bool {
			return intervals[i][0] >= p[1]
		})
		if i < n {
			ans[p[2]] = intervals[i][2]
		} else {
			ans[p[2]] = -1
		}
	}
	return ans
}

//------------------------Leetcode Problem 436------------------------
/*
 * https://leetcode.cn/problems/find-right-interval/
 * 执行用时：36ms 在所有Go提交中击败了96.12%的用户
 * 占用内存：7.3MB 在所有Go提交中击败了62.30%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		n, _ := strconv.Atoi(scanner.Text())
		intervals := make([][]int, n)
		for i := range intervals {
			scanner.Scan()
			intervals[i] = stringArrayToIntArray(strings.Fields(scanner.Text()))
		}
		fmt.Printf("Output: %v\n", findRightInterval(intervals))
	}
}
