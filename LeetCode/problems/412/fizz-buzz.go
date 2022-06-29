package main

import (
	"bufio"
	. "fmt"
	"os"
	"strconv"
)

// https://leetcode.cn/problems/fizz-buzz/
//------------------------Leetcode Problem 412------------------------
func fizzBuzz(n int) []string {
	var ans []string
	for i := 1; i <= n; i++ {
		if i%3 == 0 && i%5 == 0 {
			ans = append(ans, "FizzBuzz")
		} else if i%3 == 0 {
			ans = append(ans, "Fizz")
		} else if i%5 == 0 {
			ans = append(ans, "Buzz")
		} else {
			ans = append(ans, strconv.Itoa(i))
		}
	}
	return ans
}

//------------------------Leetcode Problem 412------------------------
/*
 * https://leetcode.cn/problems/fizz-buzz/
 * 执行用时：4ms 在所有Go提交中击败了83.64%的用户
 * 占用内存：4.1MB 在所有Go提交中击败了29.37%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		n, _ := strconv.Atoi(scanner.Text())
		Printf("Output: %v\n", fizzBuzz(n))
	}
}
