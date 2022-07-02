package main

import (
	"bufio"
	. "fmt"
	"os"
	"strconv"
	"strings"
)

// https://leetcode.cn/problems/baseball-game/
//------------------------Leetcode Problem 682------------------------
func calPoints(ops []string) int {
	var points []int
	var ans int
	for _, op := range ops {
		n := len(points)
		switch op[0] {
		case '+':
			ans += points[n-1] + points[n-2]
			points = append(points, points[n-1]+points[n-2])
		case 'D':
			ans += points[n-1] * 2
			points = append(points, 2*points[n-1])
		case 'C':
			ans -= points[n-1]
			points = points[:len(points)-1]
		default:
			pt, _ := strconv.Atoi(op)
			ans += pt
			points = append(points, pt)
		}
	}
	return ans
}

//------------------------Leetcode Problem 682------------------------
/*
 * https://leetcode.cn/problems/baseball-game/
 * 执行用时：0ms 在所有Go提交中击败了95.56%的用户
 * 占用内存：2.5MB 在所有Go提交中击败了88.49%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		ops := strings.Fields(scanner.Text())
		Printf("Output: %v\n", calPoints(ops))
	}
}
