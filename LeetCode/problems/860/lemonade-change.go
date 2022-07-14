package main

import (
	"bufio"
	. "fmt"
	"github.com/Don2025/GoCode/utils"
	"os"
)

// https://leetcode.cn/problems/lemonade-change/
//------------------------Leetcode Problem 860------------------------
func lemonadeChange(bills []int) bool {
	var five, ten int
	for _, x := range bills {
		if x == 5 {
			five++
		} else if x == 10 {
			five--
			ten++
		} else if ten > 0 {
			five--
			ten--
		} else {
			five -= 3
		}
		if five < 0 {
			return false
		}
	}
	return true
}

//------------------------Leetcode Problem 860------------------------
/*
 * https://leetcode.cn/problems/lemonade-change/
 * 执行用时：92ms 在所有Go提交中击败了87.81%的用户
 * 占用内存：8.3MB 在所有Go提交中击败了73.38%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		bills := utils.StringToInts(scanner.Text())
		Printf("Output: %v\n", lemonadeChange(bills))
	}
}
