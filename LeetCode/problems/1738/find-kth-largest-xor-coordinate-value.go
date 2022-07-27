package main

import (
	"bufio"
	. "fmt"
	"github.com/Don2025/GoCode/utils"
	"os"
	"sort"
	"strconv"
)

// https://leetcode.cn/problems/find-kth-largest-xor-coordinate-value/
//------------------------Leetcode Problem 1738------------------------
func kthLargestValue(matrix [][]int, k int) int {
	n, m := len(matrix), len(matrix[0])
	location := make([]int, 0, m*n)
	a := make([][]int, n+1)
	a[0] = make([]int, m+1)
	for i, row := range matrix {
		a[i+1] = make([]int, m+1)
		for j, x := range row {
			a[i+1][j+1] = a[i+1][j] ^ a[i][j+1] ^ a[i][j] ^ x
			location = append(location, a[i+1][j+1])
		}
	}
	sort.Sort(sort.Reverse(sort.IntSlice(location)))
	return location[k-1]
}

//------------------------Leetcode Problem 1738------------------------
/*
 * https://leetcode.cn/problems/find-kth-largest-xor-coordinate-value/
 * 执行用时：556ms 在所有Go提交中击败了33.72%的用户
 * 占用内存：34.9MB 在所有Go提交中击败了59.35%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		n, _ := strconv.Atoi(scanner.Text())
		matrix := make([][]int, n)
		for i := 0; i < n; i++ {
			scanner.Scan()
			matrix[i] = utils.StringToInts(scanner.Text())
		}
		scanner.Scan()
		k, _ := strconv.Atoi(scanner.Text())
		Printf("Output: %v\n", kthLargestValue(matrix, k))
	}
}
