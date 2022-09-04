package main

import (
	"bufio"
	. "fmt"
	"github.com/Don2025/GoCode/utils"
	"os"
	"strconv"
)

// https://leetcode.cn/problems/special-positions-in-a-binary-matrix/
// ------------------------Leetcode Problem 1582------------------------
func numSpecial(mat [][]int) int {
	rowsSum := make([]int, len(mat))
	colsSum := make([]int, len(mat[0]))
	var ans int
	for i := range mat {
		for j := range mat[i] {
			rowsSum[i] += mat[i][j]
			colsSum[j] += mat[i][j]
		}
	}
	for i := range mat {
		for j := range mat[i] {
			if mat[i][j] == 1 && rowsSum[i] == 1 && colsSum[j] == 1 {
				ans++
			}
		}
	}
	return ans
}

//------------------------Leetcode Problem 1582------------------------
/*
 * https://leetcode.cn/problems/special-positions-in-a-binary-matrix/
 * 执行用时：16ms 在所有Go提交中击败了87.50%的用户
 * 占用内存：6.3MB 在所有Go提交中击败了81.25%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		n, _ := strconv.Atoi(scanner.Text())
		mat := make([][]int, n)
		for i := range mat {
			scanner.Scan()
			mat[i] = utils.StringToInts(scanner.Text())
		}
		Printf("Output: %v\n", numSpecial(mat))
	}
}
