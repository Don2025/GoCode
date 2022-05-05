package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func imageSmoother(img [][]int) [][]int {
	row, col := len(img), len(img[0])
	ans := make([][]int, row)
	for i := range ans {
		ans[i] = make([]int, col)
		for j := range ans[i] {
			var sum, cnt int
			for _, x := range img[max(i-1, 0):min(i+2, row)] {
				for _, y := range x[max(j-1, 0):min(j+2, col)] {
					sum += y
					cnt++
				}
			}
			ans[i][j] = sum / cnt
		}
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		n, _ := strconv.Atoi(input.Text())
		img := make([][]int, n)
		for i := 0; i < n; i++ {
			input.Scan()
			img[i] = stringArrayToIntArray(strings.Fields(input.Text()))
		}
		ans := imageSmoother(img)
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
 * 执行用时：36ms 在所有Go提交中击败了93.33%的用户
 * 占用内存：7.1MB 在所有Go提交中击败了84.44%的用户
**/
