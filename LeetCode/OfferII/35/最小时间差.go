package main

import (
	"bufio"
	. "fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

// https://leetcode.cn/problems/569nqc/
// ------------------------剑指 Offer II Problem 35------------------------
func findMinDifference(timePoints []string) int {
	ans := 24 * 60
	if len(timePoints) > ans {
		return 0
	}
	var mins []int
	for _, time := range timePoints {
		t := strings.Split(time, ":")
		h, _ := strconv.Atoi(t[0])
		m, _ := strconv.Atoi(t[1])
		mins = append(mins, h*60+m)
	}
	sort.Ints(mins)
	mins = append(mins, mins[0]+ans) // 用于处理最大值和最小值的差值这种特殊情况
	for i := 1; i < len(mins); i++ {
		ans = min(ans, mins[i]-mins[i-1])
	}
	return ans
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// ------------------------剑指 Offer II Problem 35------------------------
/*
 * https://leetcode.cn/problems/569nqc/
 * 执行用时：4ms 在所有Go提交中击败了94.96%的用户
 * 占用内存：4.2MB 在所有Go提交中击败了52.16%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		timePoints := strings.Fields(scanner.Text())
		Printf("Output: %v\n", findMinDifference(timePoints))
	}
}
