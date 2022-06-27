package main

import (
	"bufio"
	. "fmt"
	"os"
	"strconv"
)

// https://leetcode.cn/problems/excel-sheet-column-title/
//------------------------Leetcode Problem 168------------------------
func convertToTitle(columnNumber int) string {
	var ans string
	if columnNumber <= 0 {
		return ans
	}
	for columnNumber > 0 {
		columnNumber--
		ans = string(columnNumber%26+'A') + ans
		columnNumber /= 26
	}
	return ans
}

//------------------------Leetcode Problem 168------------------------
/*
 * https://leetcode.cn/problems/excel-sheet-column-title/
 * 执行用时：0ms 在所有Go提交中击败了100.00%的用户
 * 占用内存：1.9MB 在所有Go提交中击败了99.48%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		n, _ := strconv.Atoi(scanner.Text())
		Printf("Output: %v\n", convertToTitle(n))
	}
}
