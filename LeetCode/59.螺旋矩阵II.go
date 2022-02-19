package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func generateMatrix(n int) [][]int {
	matrix := make([][]int, n)
	for i := range matrix {
		matrix[i] = make([]int, n)
	}
	num := 1
	left, right, top, bottom := 0, n-1, 0, n-1
	for left <= right && top <= bottom {
		for column := left; column <= right; column++ {
			matrix[top][column] = num
			num++
		}
		for row := top + 1; row <= bottom; row++ {
			matrix[row][right] = num
			num++
		}
		if left < right && top < bottom {
			for column := right - 1; column > left; column-- {
				matrix[bottom][column] = num
				num++
			}
			for row := bottom; row > top; row-- {
				matrix[row][left] = num
				num++
			}
		}
		left++
		right--
		top++
		bottom--
	}
	return matrix
}

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		n, _ := strconv.Atoi(input.Text())
		ans := generateMatrix(n)
		for _, row := range ans {
			fmt.Printf("%v ", row)
		}
		fmt.Println()
	}
}

/*
 * 执行用时：0ms 在所有Go提交中击败了100.00%的用户
 * 占用内存：2MB 在所有Go提交中击败了42.92%的用户
**/
