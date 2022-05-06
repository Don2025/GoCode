package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func countNegatives(grid [][]int) int {
	var cnt int
	for _, row := range grid {
		l, r, pos := 0, len(row)-1, len(row)
		cnt += pos
		for l <= r {
			mid := l + (r-l)>>1
			if row[mid] < 0 {
				pos = mid
				r = mid - 1
			} else {
				l = mid + 1
			}
		}
		cnt -= pos
	}
	return cnt
}

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		n, _ := strconv.Atoi(input.Text())
		input.Scan()
		grid := make([][]int, n)
		for i := range grid {
			input.Scan()
			grid[i] = stringArrayToIntArray(strings.Fields(input.Text()))
		}
		println(countNegatives(grid))
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
 * 执行用时：12ms 在所有Go提交中击败了89.63%的用户
 * 占用内存：6.2MB 在所有Go提交中击败了100.00%的用户
**/
