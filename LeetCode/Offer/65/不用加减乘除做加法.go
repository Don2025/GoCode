package main

import (
	"bufio"
	. "fmt"
	"os"
)

// https://leetcode.cn/problems/bu-yong-jia-jian-cheng-chu-zuo-jia-fa-lcof/
// ------------------------剑指 Offer I Problem 65------------------------
func add(a int, b int) int {
	var c int
	for b != 0 { // 进位不等于0
		c = (a & b) << 1 // 算出进位
		a ^= b           // 算出不带进位的和
		b = c            // 更新进位
	}
	return a
}

// ------------------------剑指 Offer I Problem 65------------------------
/*
 * https://leetcode.cn/problems/bu-yong-jia-jian-cheng-chu-zuo-jia-fa-lcof/
 * 执行用时：0ms 在所有Go提交中击败了100.00%的用户
 * 占用内存：1.8MB 在所有Go提交中击败了82.63%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		var a, b int
		Sscanf(scanner.Text(), "%d %d", &a, &b)
		Printf("Output: %d\n", add(a, b))
	}
}
