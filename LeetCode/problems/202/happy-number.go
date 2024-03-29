package main

import (
	"bufio"
	. "fmt"
	"os"
	"strconv"
)

// https://leetcode.cn/problems/happy-number/
//------------------------Leetcode Problem 202------------------------
func isHappy(n int) bool {
	var step func(int) int
	step = func(n int) int {
		var sum int
		for n > 0 {
			t := n % 10
			sum += t * t
			n /= 10
		}
		return sum
	}
	slow, fast := n, step(n)
	for slow != fast && fast != 1 {
		slow = step(slow)
		fast = step(step(fast))
	}
	return fast == 1
}

//------------------------Leetcode Problem 202------------------------
/*
 * https://leetcode.cn/problems/happy-number/
 * 执行用时：0ms 在所有Go提交中击败了100.00%的用户
 * 占用内存：1.9MB 在所有Go提交中击败了89.21%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		n, _ := strconv.Atoi(scanner.Text())
		Printf("Output: %v\n", isHappy(n))
	}
}
