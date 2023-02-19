package main

import (
	"bufio"
	. "fmt"
	"os"
)

// https://leetcode.cn/problems/determine-color-of-a-chessboard-square/
//------------------------Leetcode Problem 1812------------------------
func squareIsWhite(coordinates string) bool {
	return ((coordinates[0]-'a'+10)^(coordinates[1])-'0')&1 == 0
}

//------------------------Leetcode Problem 1812------------------------
/*
 * https://leetcode.cn/problems/determine-color-of-a-chessboard-square/
 * 执行用时：4ms 在所有Go提交中击败了12.00%的用户
 * 占用内存：1.8MB 在所有Go提交中击败了80.00%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		Printf("Output: %v\n", squareIsWhite(scanner.Text()))
	}
}
