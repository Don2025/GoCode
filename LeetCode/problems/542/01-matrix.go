package main

import (
	"bufio"
	. "fmt"
	"github.com/Don2025/GoCode/utils"
	"math"
	"os"
	"strconv"
)

// https://leetcode.cn/problems/01-matrix/
//------------------------Leetcode Problem 542------------------------
func updateMatrix(mat [][]int) [][]int {
	row, col := len(mat), len(mat[0])
	queue := make([][]int, 0)
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if mat[i][j] == 0 {
				queue = append(queue, []int{i, j})
			} else {
				mat[i][j] = -math.MaxInt32
			}
		}
	}
	direction := [][]int{{0, -1}, {0, 1}, {-1, 0}, {1, 0}} //上下左右
	for len(queue) != 0 {
		point := queue[0]
		queue = queue[1:]
		for _, v := range direction {
			x, y := point[0]+v[0], point[1]+v[1]
			if x >= 0 && x < row && y >= 0 && y < col && mat[x][y] == -math.MaxInt32 {
				mat[x][y] = mat[point[0]][point[1]] + 1
				queue = append(queue, []int{x, y})
			}
		}
	}
	return mat
}

//------------------------Leetcode Problem 542------------------------
/*
 * https://leetcode.cn/problems/01-matrix/
 * 执行用时：112ms 在所有Go提交中击败了6.25%的用户
 * 占用内存：7.8MB 在所有Go提交中击败了58.33%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		n, _ := strconv.Atoi(scanner.Text())
		arr := make([][]int, n)
		for i := 0; i < n; i++ {
			scanner.Scan()
			arr[i] = utils.StringToInts(scanner.Text())
		}
		Printf("Output: %v\n", updateMatrix(arr))
	}
}
