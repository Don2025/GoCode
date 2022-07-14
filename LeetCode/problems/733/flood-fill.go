package main

import (
	"bufio"
	. "fmt"
	"github.com/Don2025/GoCode/utils"
	"os"
	"strconv"
)

// https://leetcode.cn/problems/flood-fill/
//------------------------Leetcode Problem 733------------------------
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

//------------------------Leetcode Problem 733------------------------
/*
 * https://leetcode.cn/problems/flood-fill/
 * 执行用时：4ms 在所有Go提交中击败了96.82%的用户
 * 占用内存：4.1MB 在所有Go提交中击败了75.66%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		n, _ := strconv.Atoi(scanner.Text())
		image := make([][]int, n)
		for i := 0; i < n; i++ {
			scanner.Scan()
			image[i] = utils.StringToInts(scanner.Text())
		}
		scanner.Scan()
		arr := utils.StringToInts(scanner.Text())
		sr, sc, newColor := arr[0], arr[1], arr[2]
		Printf("Output: %v\n", floodFill(image, sr, sc, newColor))
	}
}
