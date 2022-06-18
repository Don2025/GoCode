package main

import (
	"bufio"
	"fmt"
	"github.com/Don2025/GoCode/utils"
	"os"
	"sort"
	"strconv"
)

func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	var ans [][]int
	prev := intervals[0]
	for i := 1; i < len(intervals); i++ {
		cur := intervals[i]
		if prev[1] < cur[0] {
			ans = append(ans, prev)
			prev = cur
		} else {
			prev[1] = max(prev[1], cur[1])
		}
	}
	ans = append(ans, prev)
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		n, _ := strconv.Atoi(scanner.Text())
		intervals := make([][]int, n)
		for i := range intervals {
			scanner.Scan()
			intervals[i] = utils.StringToInts(scanner.Text())
		}
		ans := merge(intervals)
		for _, row := range ans {
			fmt.Printf("%v ", row)
		}
		fmt.Println()
	}
}

/*
 * 执行用时：16ms 在所有Go提交中击败了87.17%的用户
 * 占用内存：6.7MB 在所有Go提交中击败了66.54%的用户
**/
