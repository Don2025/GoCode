package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

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

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		n, _ := strconv.Atoi(input.Text())
		arr := make([][]int, n)
		for i := 0; i < n; i++ {
			input.Scan()
			arr[i] = stringArrayToIntArray(strings.Fields(input.Text()))
		}
		ans := updateMatrix(arr)
		for i := 0; i < len(ans); i++ {
			for j := 0; j < len(ans[0]); j++ {
				fmt.Printf("%d ", ans[i][j])
			}
			println()
		}
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
 * 执行用时：112ms 在所有Go提交中击败了6.25%的用户
 * 占用内存：7.8MB 在所有Go提交中击败了58.33%的用户
**/
