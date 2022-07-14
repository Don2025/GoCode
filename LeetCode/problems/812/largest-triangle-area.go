package main

import (
	"bufio"
	. "fmt"
	"github.com/Don2025/GoCode/utils"
	"math"
	"os"
	"strconv"
)

// https://leetcode.cn/problems/largest-triangle-area/
//------------------------Leetcode Problem 812------------------------
func largestTriangleArea(points [][]int) float64 {
	var ans float64
	triangleArea := func(x1, y1, x2, y2, x3, y3 int) float64 {
		return math.Abs(float64(x1*y2+x2*y3+x3*y1-x1*y3-x2*y1-x3*y2)) / 2
	}
	for i, p := range points {
		for j, q := range points[:i] {
			for _, r := range points[:j] {
				ans = math.Max(ans, triangleArea(p[0], p[1], q[0], q[1], r[0], r[1]))
			}
		}
	}
	return ans
}

//------------------------Leetcode Problem 812------------------------
/*
 * https://leetcode.cn/problems/largest-triangle-area/
 * 执行用时：0ms 在所有Go提交中击败了100.00%的用户
 * 占用内存：2MB 在所有Go提交中击败了100.00%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		n, _ := strconv.Atoi(scanner.Text())
		points := make([][]int, n)
		for i := range points {
			scanner.Scan()
			points[i] = utils.StringToInts(scanner.Text())
		}
		Printf("Output: %.2f\n", largestTriangleArea(points))
	}
}
