package main

import (
	"bufio"
	. "fmt"
	"os"
	"strconv"
)

// https://leetcode.cn/problems/minimum-moves-to-reach-target-score/
//------------------------Leetcode Problem 2139------------------------
func minMoves(target int, maxDoubles int) int {
	var cnt int
	for target > 1 && maxDoubles > 0 {
		if target&1 == 0 {
			target >>= 1
			maxDoubles--
		} else {
			target--
		}
		cnt++
	}
	cnt += target - 1
	return cnt
}

//------------------------Leetcode Problem 2139------------------------
/*
 * https://leetcode.cn/problems/minimum-moves-to-reach-target-score/
 * 执行用时：0ms 在所有Go提交中击败了100.00%的用户
 * 占用内存：1.9MB 在所有Go提交中击败了100.00%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		target, _ := strconv.Atoi(scanner.Text())
		scanner.Scan()
		maxDoubles, _ := strconv.Atoi(scanner.Text())
		Printf("Output: %v\n", minMoves(target, maxDoubles))
	}
}
