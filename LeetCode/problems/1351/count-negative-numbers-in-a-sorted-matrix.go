package main

import (
	"bufio"
	. "fmt"
	"github.com/Don2025/GoCode/utils"
	"os"
	"strconv"
)

// https://leetcode.cn/problems/count-negative-numbers-in-a-sorted-matrix/
// -------------------------Leetcode Problem 1351------------------------
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

// -------------------------Leetcode Problem 1351------------------------
/*
 * https://leetcode.cn/problems/count-negative-numbers-in-a-sorted-matrix/
 * 执行用时：12ms 在所有Go提交中击败了89.63%的用户
 * 占用内存：6.2MB 在所有Go提交中击败了100.00%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		n, _ := strconv.Atoi(scanner.Text())
		scanner.Scan()
		grid := make([][]int, n)
		for i := range grid {
			scanner.Scan()
			grid[i] = utils.StringToInts(scanner.Text())
		}
		Printf("Output: %v\n", countNegatives(grid))
	}
}
