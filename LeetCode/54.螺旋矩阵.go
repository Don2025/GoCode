package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func spiralOrder(matrix [][]int) []int {
	var ans []int
	row, col := len(matrix), len(matrix[0])
	if row == 0 || col == 0 {
		return ans
	}
	up, down, left, right := 0, row-1, 0, col-1
	for {
		for i := left; i <= right; i++ {
			ans = append(ans, matrix[up][i])
		}
		if up++; up > down {
			break
		}
		for i := up; i <= down; i++ {
			ans = append(ans, matrix[i][right])
		}
		if right--; right < left {
			break
		}
		for i := right; i >= left; i-- {
			ans = append(ans, matrix[down][i])
		}
		if down--; down < up {
			break
		}
		for i := down; i >= up; i-- {
			ans = append(ans, matrix[i][left])
		}
		if left++; left > right {
			break
		}
	}
	return ans
}

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		n, _ := strconv.Atoi(input.Text())
		matrix := make([][]int, n)
		for i := range matrix {
			input.Scan()
			matrix[i] = stringArrayToIntArray(strings.Fields(input.Text()))
		}
		ans := spiralOrder(matrix)
		fmt.Printf("%v\n", ans)
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
 * 占用内存：1.8MB 在所有Go提交中击败了91.19%的用户
**/
