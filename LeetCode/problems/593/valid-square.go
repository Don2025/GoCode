package main

import (
	"bufio"
	. "fmt"
	"github.com/Don2025/GoCode/utils"
	"os"
)

// https://leetcode.cn/problems/valid-square/
//------------------------Leetcode Problem 593------------------------
func validSquare(p1, p2, p3, p4 []int) bool {
	if p1[0] == p2[0] && p1[1] == p2[1] {
		return false
	}
	if help(p1, p2, p3, p4) {
		return true
	}
	if p1[0] == p3[0] && p1[1] == p3[1] {
		return false
	}
	if help(p1, p3, p2, p4) {
		return true
	}
	if p1[0] == p4[0] && p1[1] == p4[1] {
		return false
	}
	if help(p1, p4, p2, p3) {
		return true
	}
	return false
}

func pow(x int) int {
	return x * x
}

func checkLength(p1, p2 []int) bool {
	return pow(p1[0])+pow(p1[1]) == pow(p2[0])+pow(p2[1])
}

func checkMidPoint(p1, p2, p3, p4 []int) bool {
	return p1[0]+p2[0] == p3[0]+p4[0] && p1[1]+p2[1] == p3[1]+p4[1]
}

func calCos(p1, p2 []int) int {
	return p1[0]*p2[0] + p1[1]*p2[1]
}

func help(p1, p2, p3, p4 []int) bool {
	v1 := []int{p1[0] - p2[0], p1[1] - p2[1]}
	v2 := []int{p3[0] - p4[0], p3[1] - p4[1]}
	return checkLength(v1, v2) && checkMidPoint(p1, p2, p3, p4) && calCos(v1, v2) == 0
}

//------------------------Leetcode Problem 593------------------------
/*
 * https://leetcode.cn/problems/valid-square/
 * 执行用时：0ms 在所有Go提交中击败了100.00%的用户
 * 占用内存：2.1MB 在所有Go提交中击败了100.00%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		p1 := utils.StringToInts(scanner.Text())
		scanner.Scan()
		p2 := utils.StringToInts(scanner.Text())
		scanner.Scan()
		p3 := utils.StringToInts(scanner.Text())
		scanner.Scan()
		p4 := utils.StringToInts(scanner.Text())
		Printf("Output: %v\n", validSquare(p1, p2, p3, p4))
	}
}
