package main

import (
	"bufio"
	. "fmt"
	"github.com/Don2025/GoCode/utils"
	"math"
	"os"
	"strconv"
)

// https://leetcode.cn/problems/find-nearest-point-that-has-the-same-x-or-y-coordinate/
//------------------------Leetcode Problem 1779------------------------
func nearestValidPoint(x int, y int, points [][]int) int {
	ans := -1
	for i, min := 0, math.MaxInt32; i < len(points); i++ {
		dx, dy := x-points[i][0], y-points[i][1]
		if t := abs(dy + dx); dx*dy == 0 && t < min {
			min = t
			ans = i
		}
	}
	return ans
}

func abs(n int) int {
	if n < 0 {
		n = -n
	}
	return n
}

//------------------------Leetcode Problem 1779------------------------
/*
 * https://leetcode.cn/problems/find-nearest-point-that-has-the-same-x-or-y-coordinate/
 * 执行用时：88ms 在所有Go提交中击败了84.91%的用户
 * 占用内存：7.1MB 在所有Go提交中击败了72.64%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		x, _ := strconv.Atoi(scanner.Text())
		scanner.Scan()
		y, _ := strconv.Atoi(scanner.Text())
		scanner.Scan()
		n, _ := strconv.Atoi(scanner.Text())
		points := make([][]int, n)
		for i := 0; i < n; i++ {
			scanner.Scan()
			points[i] = utils.StringToInts(scanner.Text())
		}
		Printf("Output: %d\n", nearestValidPoint(x, y, points))
	}
}
