package main

import (
	"bufio"
	. "fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

// https://leetcode.cn/problems/minimum-time-difference/
//------------------------Leetcode Problem 539------------------------
func findMinDifference(timePoints []string) int {
	n := len(timePoints)
	if n > 1440 {
		return 0
	}
	cache := make([]int, n)
	for i := range timePoints {
		ss := strings.Split(timePoints[i], ":")
		hour, _ := strconv.Atoi(ss[0])
		minute, _ := strconv.Atoi(ss[1])
		cache[i] = 60*hour + minute
	}
	sort.Ints(cache)
	ans := math.MaxInt32
	for i := 1; i < n; i++ {
		ans = min(ans, cache[i]-cache[i-1])
	}
	ans = min(ans, cache[0]+1440-cache[n-1])
	return ans
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

//------------------------Leetcode Problem 539------------------------
/*
 * https://leetcode.cn/problems/minimum-time-difference/
 * 执行用时：4ms 在所有Go提交中击败了95.52%的用户
 * 占用内存：4.3MB 在所有Go提交中击败了70.15%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		timePoints := strings.Fields(scanner.Text())
		Printf("Output: %v\n", findMinDifference(timePoints))
	}
}
