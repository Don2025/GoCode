package main

import (
	"bufio"
	. "fmt"
	"os"
)

// https://leetcode.cn/problems/remove-colored-pieces-if-both-neighbors-are-the-same-color/
//------------------------Leetcode Problem 2038------------------------
func winnerOfGame(colors string) bool {
	var cnt int
	for i := 0; i < len(colors)-2; i++ {
		if colors[i:i+3] == "AAA" {
			cnt++
		} else if colors[i:i+3] == "BBB" {
			cnt--
		}
	}
	return cnt > 0
}

//------------------------Leetcode Problem 2038------------------------
/*
 * https://leetcode.cn/problems/remove-colored-pieces-if-both-neighbors-are-the-same-color/
 * 执行用时：8ms 在所有Go提交中击败了100.00%的用户
 * 占用内存：6.2MB 在所有Go提交中击败了72.73%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		Printf("Output: %v\n", winnerOfGame(scanner.Text()))
	}
}
