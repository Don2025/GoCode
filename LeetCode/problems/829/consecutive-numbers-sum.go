package main

import (
	"bufio"
	. "fmt"
	"os"
	"strconv"
)

// https://leetcode.cn/problems/consecutive-numbers-sum/
//------------------------Leetcode Problem 829------------------------
func consecutiveNumbersSum(n int) int {
	var ans int
	for i := 1; n > 0; i++ {
		if n%i == 0 {
			ans++
		}
		n -= i
	}
	return ans
}

//------------------------Leetcode Problem 829------------------------
/*
 * https://leetcode.cn/problems/consecutive-numbers-sum/
 * 执行用时：0ms 在所有Go提交中击败了100.00%的用户
 * 占用内存：1.8MB 在所有Go提交中击败了78.26%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		n, _ := strconv.Atoi(scanner.Text())
		Printf("Output: %v\n", consecutiveNumbersSum(n))
	}
}
