package main

import (
	"bufio"
	. "fmt"
	"os"
)

// https://leetcode.cn/problems/excel-sheet-column-number/
//------------------------Leetcode Problem 171------------------------
func titleToNumber(columnTitle string) int {
	var ans int
	for _, x := range columnTitle {
		ans *= 26
		ans += int(x - 'A' + 1)
	}
	return ans
}

//------------------------Leetcode Problem 171------------------------
/*
 * https://leetcode.cn/problems/excel-sheet-column-number/
 * 执行用时：0ms 在所有Go提交中击败了100.00%的用户
 * 占用内存：2MB 在所有Go提交中击败了99.75%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		Printf("Output: %v\n", titleToNumber(scanner.Text()))
	}
}
