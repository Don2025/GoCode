package main

import (
	"bufio"
	. "fmt"
	"github.com/Don2025/GoCode/utils"
	"os"
	"strconv"
)

// https://leetcode.cn/problems/image-smoother/
//------------------------Leetcode Problem 661------------------------
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

//------------------------Leetcode Problem 661------------------------
/*
 * https://leetcode.cn/problems/image-smoother/
 * 执行用时：36ms 在所有Go提交中击败了93.33%的用户
 * 占用内存：7.1MB 在所有Go提交中击败了84.44%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		n, _ := strconv.Atoi(scanner.Text())
		img := make([][]int, n)
		for i := 0; i < n; i++ {
			scanner.Scan()
			img[i] = utils.StringToInts(scanner.Text())
		}
		Printf("Output: %v\n", imageSmoother(img))
	}
}
