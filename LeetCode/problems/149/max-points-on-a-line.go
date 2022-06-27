package main

import (
	"bufio"
	. "fmt"
	"github.com/Don2025/GoCode/utils"
	"os"
	"strconv"
)

// https://leetcode.cn/problems/max-points-on-a-line/
//------------------------Leetcode Problem 149------------------------
func maxPoints(points [][]int) int {
	ans := 1
	for i := 0; i < len(points); i++ {
		x := points[i]
		for j := i + 1; j < len(points); j++ {
			y := points[j]
			cnt := 2
			for k := j + 1; k < len(points); k++ {
				z := points[k]
				if (x[0]-y[0])*(y[1]-z[1]) == (x[1]-y[1])*(y[0]-z[0]) {
					cnt++
				}
			}
			ans = max(ans, cnt)
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

//------------------------Leetcode Problem 149------------------------
/*
 * https://leetcode.cn/problems/max-points-on-a-line/
 * 执行用时：4ms 在所有Go提交中击败了88.42%的用户
 * 占用内存：2MB 在所有Go提交中击败了98.95%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		n, _ := strconv.Atoi(scanner.Text())
		points := make([][]int, n)
		for i := 0; i < n; i++ {
			scanner.Scan()
			points[i] = utils.StringToInts(scanner.Text())
		}
		Printf("Output: %v\n", maxPoints(points))
	}
}
