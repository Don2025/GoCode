package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func matrixReshape(mat [][]int, r int, c int) [][]int {
	m, n := len(mat), len(mat[0])
	if m*n != r*c {
		return mat
	}
	ans := make([][]int, r)
	for i := range ans {
		ans[i] = make([]int, c)
	}
	var x, y int
	for i := range mat {
		for j := 0; j < n; j++ {
			ans[x][y] = mat[i][j]
			y++
			if y == c {
				x++
				y = 0
			}
		}
	}
	return ans
}

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		n, _ := strconv.Atoi(input.Text())
		mat := make([][]int, n)
		for i := range mat {
			input.Scan()
			mat[i] = stringArrayToIntArray(strings.Fields(input.Text()))
		}
		input.Scan()
		arr := stringArrayToIntArray(strings.Fields(input.Text()))
		r, c := arr[0], arr[1]
		ans := matrixReshape(mat, r, c)
		for _, row := range ans {
			fmt.Printf("%v ", row)
		}
		fmt.Println()
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
 * 执行用时：12ms 在所有Go提交中击败了52.25%的用户
 * 占用内存：6.2MB 在所有Go提交中击败了81.50%的用户
**/
