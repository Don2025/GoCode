package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

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

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		ops := strings.Fields(input.Text())
		println(calPoints(ops))
	}
}

/*
 * 执行用时：0ms 在所有Go提交中击败了95.56%的用户
 * 占用内存：2.5MB 在所有Go提交中击败了88.49%的用户
**/
