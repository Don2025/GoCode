package main

import (
	"bufio"
	. "fmt"
	"github.com/Don2025/GoCode/utils"
	"os"
	"strconv"
)

// https://leetcode.cn/problems/matrix-diagonal-sum/
//------------------------Leetcode Problem 1572------------------------
func diagonalSum(mat [][]int) int {
	n := len(mat)
	sum, mid := 0, n>>1
	for i := 0; i < n; i++ {
		sum += mat[i][i] + mat[i][n-i-1]
	}
	sum -= mat[mid][mid] * (n & 1)
	return sum
}

//------------------------Leetcode Problem 1572------------------------
/*
 * https://leetcode.cn/problems/matrix-diagonal-sum/
 * 执行用时：12ms 在所有Go提交中击败了84.62%的用户
 * 占用内存：5.3MB 在所有Go提交中击败了100.00%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		n, _ := strconv.Atoi(scanner.Text())
		mat := make([][]int, n)
		for i := 0; i < n; i++ {
			scanner.Scan()
			mat[i] = utils.StringToInts(scanner.Text())
		}
		Printf("Output: %v\n", diagonalSum(mat))
	}
}
