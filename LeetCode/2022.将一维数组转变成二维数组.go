package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

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

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		original := stringArrayToIntArray(strings.Fields(input.Text()))
		input.Scan()
		m, _ := strconv.Atoi(input.Text())
		input.Scan()
		n, _ := strconv.Atoi(input.Text())
		ans := construct2DArray(original, m, n)
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
 * 执行用时：108ms 在所有Go提交中击败了80.10%的用户
 * 占用内存：9MB 在所有Go提交中击败了41.29%的用户
**/
