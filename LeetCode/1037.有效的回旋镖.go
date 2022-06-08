package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func isBoomerang(points [][]int) bool {
	x1, y1 := points[1][0]-points[0][0], points[1][1]-points[0][1]
	x2, y2 := points[2][0]-points[0][0], points[2][1]-points[0][1]
	return x1*y2 != x2*y1
}

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		n, _ := strconv.Atoi(input.Text())
		points := make([][]int, n)
		for i := range points {
			input.Scan()
			points[i] = stringArrayToIntArray(strings.Fields(input.Text()))
		}
		println(isBoomerang(points))
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
 * 执行用时：0ms 在所有Go提交中击败了100.00%的用户
 * 占用内存：1.9MB 在所有Go提交中击败了100.00%的用户
**/
