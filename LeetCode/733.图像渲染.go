package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func floodFill(image [][]int, sr int, sc int, newColor int) [][]int {
	var dfs func([][]int, int, int, int, int)
	dfs = func(g [][]int, i int, j int, c int, nc int) {
		if i < 0 || i >= len(image) || j < 0 || j >= len(image[0]) {
			return
		}
		if g[i][j] == c {
			g[i][j] = nc
			dfs(g, i-1, j, c, nc)
			dfs(g, i+1, j, c, nc)
			dfs(g, i, j-1, c, nc)
			dfs(g, i, j+1, c, nc)
		}
	}
	if image[sr][sc] != newColor {
		dfs(image, sr, sc, image[sr][sc], newColor)
	}
	return image
}

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		n, _ := strconv.Atoi(input.Text())
		image := make([][]int, n)
		for i := 0; i < n; i++ {
			input.Scan()
			image[i] = stringArrayToIntArray(strings.Fields(input.Text()))
		}
		input.Scan()
		arr := stringArrayToIntArray(strings.Fields(input.Text()))
		sr, sc, newColor := arr[0], arr[1], arr[2]
		ans := floodFill(image, sr, sc, newColor)
		for _, row := range ans {
			for _, x := range row {
				fmt.Printf("%d ", x)
			}
			fmt.Println()
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
 * 执行用时：4ms 在所有Go提交中击败了96.82%的用户
 * 占用内存：4.1MB 在所有Go提交中击败了75.66%的用户
**/
