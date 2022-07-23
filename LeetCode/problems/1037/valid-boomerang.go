package main

import (
	"bufio"
	. "fmt"
	"github.com/Don2025/GoCode/utils"
	"os"
	"strconv"
)

// https://leetcode.cn/problems/valid-boomerang/
//------------------------Leetcode Problem 1037------------------------
func isBoomerang(points [][]int) bool {
	x1, y1 := points[1][0]-points[0][0], points[1][1]-points[0][1]
	x2, y2 := points[2][0]-points[0][0], points[2][1]-points[0][1]
	return x1*y2 != x2*y1
}

//------------------------Leetcode Problem 1037------------------------
/*
 * https://leetcode.cn/problems/valid-boomerang/
 * 执行用时：0ms 在所有Go提交中击败了100.00%的用户
 * 占用内存：1.9MB 在所有Go提交中击败了100.00%的用户
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
		Printf("Output: %v\n", isBoomerang(points))
	}
}
