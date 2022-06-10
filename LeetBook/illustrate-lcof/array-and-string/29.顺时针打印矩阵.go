package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 {
		return nil
	}
	left, top, right, bottom := 0, 0, len(matrix[0])-1, len(matrix)-1
	var ans []int
	for {
		// left->right
		for i := left; i <= right; i++ {
			ans = append(ans, matrix[top][i])
		}
		top++
		if top > bottom {
			break
		}
		// top->bottom
		for i := top; i <= bottom; i++ {
			ans = append(ans, matrix[i][right])
		}
		right--
		if left > right {
			break
		}
		// right->left
		for i := right; i >= left; i-- {
			ans = append(ans, matrix[bottom][i])
		}
		bottom--
		if top > bottom {
			break
		}
		// bottom->top
		for i := bottom; i >= top; i-- {
			ans = append(ans, matrix[i][left])
		}
		left++
		if left > right {
			break
		}
	}
	return ans
}

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		n, _ := strconv.Atoi(input.Text()) //行数
		matrix := make([][]int, n)
		for i := 0; i < n; i++ {
			input.Scan()
			matrix[i] = stringArrayToIntArray(strings.Fields(input.Text()))
		}
		ans := spiralOrder(matrix)
		for _, x := range ans {
			fmt.Printf("%d ", x)
		}
		println()
	}
}

/*
 * 执行用时：16ms 在所有Go提交中击败了50.22%的用户
 * 占用内存：6.5MB 在所有Go提交中击败了46.17%的用户
**/
