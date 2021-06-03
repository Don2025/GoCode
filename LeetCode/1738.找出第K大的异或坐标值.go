package main

import (
	"bufio"
	"os"
	"sort"
	"strconv"
	"strings"
)

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

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		n, _ := strconv.Atoi(input.Text())
		matrix := make([][]int, n)
		for i := 0; i < n; i++ {
			input.Scan()
			matrix[i] = stringArrayToIntArray(strings.Fields(input.Text()))
		}
		input.Scan()
		k, _ := strconv.Atoi(input.Text())
		println(kthLargestValue(matrix, k))
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
 * 执行用时：556ms 在所有Go提交中击败了33.72%的用户
 * 占用内存：34.9MB 在所有Go提交中击败了59.35%的用户
**/
