package main

import (
	"bufio"
	"math"
	"os"
	"strconv"
	"strings"
)

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

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		x, _ := strconv.Atoi(input.Text())
		input.Scan()
		y, _ := strconv.Atoi(input.Text())
		input.Scan()
		n, _ := strconv.Atoi(input.Text())
		points := make([][]int, n)
		for i := 0; i < n; i++ {
			input.Scan()
			points[i] = stringArrayToIntArray(strings.Fields(input.Text()))
		}
		println(nearestValidPoint(x, y, points))
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
 * 执行用时：88ms 在所有Go提交中击败了84.91%的用户
 * 占用内存：7.1MB 在所有Go提交中击败了72.64%的用户
**/
