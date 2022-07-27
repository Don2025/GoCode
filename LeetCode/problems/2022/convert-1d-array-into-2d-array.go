package main

import (
	"bufio"
	. "fmt"
	"github.com/Don2025/GoCode/utils"
	"os"
	"strconv"
)

// https://leetcode.cn/problems/convert-1d-array-into-2d-array/
//------------------------Leetcode Problem 2022------------------------
func construct2DArray(original []int, m int, n int) [][]int {
	if m*n != len(original) {
		return [][]int{}
	}
	ans := make([][]int, m)
	var idx int
	for i := 0; i < m; i++ {
		ans[i] = make([]int, n)
		for j := 0; j < n; j++ {
			ans[i][j] = original[idx]
			idx++
		}
	}
	return ans
}

//------------------------Leetcode Problem 2022------------------------
/*
 * https://leetcode.cn/problems/convert-1d-array-into-2d-array/
 * 执行用时：108ms 在所有Go提交中击败了80.10%的用户
 * 占用内存：9MB 在所有Go提交中击败了41.29%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		original := utils.StringToInts(scanner.Text())
		scanner.Scan()
		m, _ := strconv.Atoi(scanner.Text())
		scanner.Scan()
		n, _ := strconv.Atoi(scanner.Text())
		Printf("Output: %v\n", construct2DArray(original, m, n))
	}
}
