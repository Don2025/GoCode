package main

import (
	"bufio"
	. "fmt"
	"github.com/Don2025/GoCode/utils"
	"os"
)

// https://leetcode.cn/problems/utf-8-validation/
//------------------------Leetcode Problem 393------------------------
func validUtf8(data []int) bool {
	var n int
	for _, x := range data {
		if n == 0 {
			if (x & 0b10000000) == 0b00000000 {
				n = 0
			} else if (x & 0b11100000) == 0b11000000 {
				n = 1
			} else if (x & 0b11110000) == 0b11100000 {
				n = 2
			} else if (x & 0b11111000) == 0b11110000 {
				n = 3
			} else {
				return false
			}
		} else {
			if (x & 0b11000000) != 0b10000000 {
				return false
			}
			n--
		}
	}
	return n == 0
}

//------------------------Leetcode Problem 393------------------------
/*
 * https://leetcode.cn/problems/utf-8-validation/
 * 执行用时：8ms 在所有Go提交中击败了97.44%的用户
 * 占用内存：5MB 在所有Go提交中击败了92.31%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		data := utils.StringToInts(scanner.Text())
		Printf("Output: %v\n", validUtf8(data))
	}
}
