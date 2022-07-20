package main

import (
	"bufio"
	. "fmt"
	"github.com/Don2025/GoCode/utils"
	"os"
	"strconv"
)

// https://leetcode.cn/problems/shift-2d-grid/
//------------------------Leetcode Problem 1260------------------------
func shiftGrid(grid [][]int, k int) [][]int {
	var arr []int
	for _, row := range grid {
		arr = append(arr, row...)
	}
	tmp := make([]int, len(arr))
	copy(tmp, arr)
	for i := range arr {
		arr[(k+i)%len(arr)] = tmp[i]
	}
	ans := make([][]int, len(grid))
	for i, row := range grid {
		for j := range row {
			ans[i] = append(ans[i], arr[i*len(grid[0])+j])
		}
	}
	return ans
}

//------------------------Leetcode Problem 1260------------------------
/*
 * https://leetcode.cn/problems/shift-2d-grid/
 * 执行用时：20ms 在所有Go提交中击败了75.61%的用户
 * 占用内存：6.9MB 在所有Go提交中击败了19.51%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		n, _ := strconv.Atoi(scanner.Text())
		grid := make([][]int, n)
		for i := range grid {
			scanner.Scan()
			grid[i] = utils.StringToInts(scanner.Text())
		}
		k, _ := strconv.Atoi(scanner.Text())
		Printf("Output: %v\n", shiftGrid(grid, k))
	}
}
