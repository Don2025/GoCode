package main

import (
	"bufio"
	. "fmt"
	"github.com/Don2025/GoCode/utils"
	"os"
)

// https://leetcode.cn/problems/number-of-lines-to-write-string/
//------------------------Leetcode Problem 806------------------------
func numberOfLines(widths []int, s string) []int {
	row, col := 1, 0
	for _, x := range s {
		col += widths[x-'a']
		if col > 100 {
			col = widths[x-'a']
			row++
		}
	}
	return []int{row, col}
}

//------------------------Leetcode Problem 806------------------------
/*
 * https://leetcode.cn/problems/number-of-lines-to-write-string/
 * 执行用时：0ms 在所有Go提交中击败了100.00%的用户
 * 占用内存：1.8MB 在所有Go提交中击败了68.18%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		widths := utils.StringToInts(scanner.Text())
		scanner.Scan()
		s := scanner.Text()
		Printf("Output: %v\n", numberOfLines(widths, s))
	}
}
