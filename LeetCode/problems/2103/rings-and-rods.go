package main

import (
	"bufio"
	. "fmt"
	"os"
)

// https://leetcode.cn/problems/rings-and-rods/
//------------------------Leetcode Problem 2103------------------------
func countPoints(rings string) int {
	a := [10][3]bool{}
	for i := 0; i < len(rings); i++ {
		if rings[i] == 'R' {
			a[rings[i+1]-'0'][0] = true
		} else if rings[i] == 'G' {
			a[rings[i+1]-'0'][1] = true
		} else if rings[i] == 'B' {
			a[rings[i+1]-'0'][2] = true
		}
	}
	var cnt int
	for i := range a {
		if a[i][0] && a[i][1] && a[i][2] {
			cnt++
		}
	}
	return cnt
}

//------------------------Leetcode Problem 2103------------------------
/*
 * https://leetcode.cn/problems/rings-and-rods/
 * 执行用时：0ms 在所有Go提交中击败了100.00%的用户
 * 占用内存：1.9MB 在所有Go提交中击败了89.08%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		Printf("Output: %d\n", countPoints(scanner.Text()))
	}
}
