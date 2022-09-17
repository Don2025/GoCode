package main

import (
	"bufio"
	. "fmt"
	"os"
)

// https://leetcode.cn/count-operations-to-obtain-zero/
//------------------------Leetcode Problem 6219------------------------
func countOperations(num1 int, num2 int) int {
	var cnt int
	for num1 > 0 && num2 > 0 {
		if num1 > num2 {
			num1 -= num2
		} else {
			num2 -= num1
		}
		cnt++
	}
	return cnt
}

//------------------------Leetcode Problem 6219------------------------
/*
 * https://leetcode.cn/count-operations-to-obtain-zero/
 * 执行用时：0ms 在所有Go提交中击败了100.00%的用户
 * 占用内存：1.8MB 在所有Go提交中击败了56.52%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		var num1, num2 int
		Sscanf(scanner.Text(), "%d %d", &num1, &num2)
		Printf("Output: %d\n", countOperations(num1, num2))
	}
}
