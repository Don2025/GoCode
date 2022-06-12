package main

import (
	"bufio"
	. "fmt"
	"github.com/Don2025/GoCode/utils"
	"os"
	"sort"
	"strconv"
)

// https://leetcode.cn/problems/SsGoHC/
// ------------------------剑指 Offer II Problem 74------------------------
func merge(intervals [][]int) (ans [][]int) {
	n := len(intervals)
	if n == 0 {
		return nil
	}
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	left, right := intervals[0][0], intervals[0][1]
	for i := 1; i < n; i++ {
		l, r := intervals[i][0], intervals[i][1]
		if l > right {
			ans = append(ans, []int{left, right})
			left, right = l, r
		} else if r > right {
			right = r
		}
	}
	ans = append(ans, []int{left, right})
	return ans
}

// ------------------------剑指 Offer II Problem 74------------------------
/*
 * https://leetcode.cn/problems/SsGoHC/
 * 执行用时：8ms 在所有Go提交中击败了88.78%的用户
 * 占用内存：4.3MB 在所有Go提交中击败了32.65%的用户
**/
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	Printf("Input the number of rows of matrix:")
	for scanner.Scan() {
		n, _ := strconv.Atoi(scanner.Text())
		intervals := make([][]int, n)
		for i := range intervals {
			Printf("Input the %d row of numbers separated by spaces:\n", i+1)
			scanner.Scan()
			intervals[i] = utils.StringToInts(scanner.Text())
		}
		Printf("Output: %v\n", merge(intervals))
		Printf("Input the number of rows of matrix:")
	}
}
