package main

import (
	"bufio"
	. "fmt"
	"os"
)

// https://leetcode.cn/problems/bulb-switcher-ii/
//------------------------Leetcode Problem 670------------------------

func flipLights(n int, presses int) int {
	if presses == 0 {
		return 1
	}
	if n == 1 {
		return 2
	} else if n == 2 {
		if presses == 1 {
			return 3
		}
		return 4
	} else {
		if presses == 1 {
			return 4
		} else if presses == 2 {
			return 7
		} else {
			return 8
		}
	}
}

//------------------------Leetcode Problem 670------------------------
/*
 * https://leetcode.cn/problems/bulb-switcher-ii/
 * 执行用时：0ms 在所有Go提交中击败了100.00%的用户
 * 占用内存：1.8MB 在所有Go提交中击败了56.52%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		var n, presses int
		Sscanf(scanner.Text(), "%d %d", &n, &presses)
		Printf("Output: %d\n", flipLights(n, presses))
	}
}
