package main

import (
	"bufio"
	. "fmt"
	"github.com/Don2025/GoCode/utils"
	"os"
)

// https://leetcode.cn/problems/chalkboard-xor-game/
//------------------------Leetcode Problem 810------------------------
func xorGame(nums []int) bool {
	var flag, n = 0, len(nums)
	for _, x := range nums {
		flag ^= x
	}
	if flag == 0 || n%2 == 0 {
		return true
	}
	return false
}

//------------------------Leetcode Problem 810------------------------
/*
 * https://leetcode.cn/problems/chalkboard-xor-game/
 * 执行用时：8ms 在所有Go提交中击败了66.67%的用户
 * 占用内存：3.9MB 在所有Go提交中击败了100.00%的用户
**/

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		a := utils.StringToInts(input.Text())
		Printf("Output: %v\n", xorGame(a))
	}
}
