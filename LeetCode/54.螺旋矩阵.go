package main

import (
	"bufio"
	"fmt"
	"github.com/Don2025/GoCode/utils"
	"os"
	"strconv"
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
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		n, _ := strconv.Atoi(scanner.Text())
		matrix := make([][]int, n)
		for i := range matrix {
			scanner.Scan()
			matrix[i] = utils.StringToInts(scanner.Text())
		}
		ans := spiralOrder(matrix)
		fmt.Printf("%v\n", ans)
	}
}

/*
 * 执行用时：0ms 在所有Go提交中击败了100.00%的用户
 * 占用内存：1.8MB 在所有Go提交中击败了91.19%的用户
**/
