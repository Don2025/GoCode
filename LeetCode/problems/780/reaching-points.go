package main

import (
	"bufio"
	. "fmt"
	"os"
)

// https://leetcode.cn/problems/reaching-points/
//------------------------Leetcode Problem 780------------------------
func reachingPoints(sx int, sy int, tx int, ty int) bool {
	for tx > 0 && ty > 0 {
		if sx == tx && sy == ty {
			return true
		}
		if tx > ty {
			tx -= ty * max((tx-sx)/ty, 1)
		} else {
			ty -= tx * max((ty-sy)/tx, 1)
		}
	}
	return false
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

//------------------------Leetcode Problem 780------------------------
/*
 * https://leetcode.cn/problems/reaching-points/
 * 执行用时：0ms 在所有Go提交中击败了100.00%的用户
 * 占用内存：1.9MB 在所有Go提交中击败了100.00%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		var sx, sy, tx, ty int
		Sscanf(scanner.Text(), "%d %d %d %d", &sx, &sy, &tx, &ty)
		Printf("Output: %v\n", reachingPoints(sx, sy, tx, ty))
	}
}
