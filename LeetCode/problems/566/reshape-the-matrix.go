package main

import (
	"bufio"
	. "fmt"
	"github.com/Don2025/GoCode/utils"
	"os"
	"strconv"
)

// https://leetcode.cn/problems/reshape-the-matrix/
//------------------------Leetcode Problem 566------------------------
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
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		n, _ := strconv.Atoi(scanner.Text())
		mat := make([][]int, n)
		for i := range mat {
			scanner.Scan()
			mat[i] = utils.StringToInts(scanner.Text())
		}
		scanner.Scan()
		arr := utils.StringToInts(scanner.Text())
		r, c := arr[0], arr[1]
		Printf("Output: %v\n", matrixReshape(mat, r, c))
	}
}

//------------------------Leetcode Problem 566------------------------
/*
 * https://leetcode.cn/problems/reshape-the-matrix/
 * 执行用时：12ms 在所有Go提交中击败了52.25%的用户
 * 占用内存：6.2MB 在所有Go提交中击败了81.50%的用户
**/
