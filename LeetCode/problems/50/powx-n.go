package main

import (
	"bufio"
	. "fmt"
	"os"
	"strconv"
)

// https://leetcode.cn/problems/powx-n/
//------------------------Leetcode Problem 50------------------------
func myPow(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	if n < 0 {
		return 1 / myPow(x, -n)
	}
	if n%2 == 1 {
		return x * myPow(x, n-1)
	}
	return myPow(x*x, n/2)
}

//------------------------Leetcode Problem 50------------------------
/*
 * https://leetcode.cn/problems/powx-n/
 * 执行用时：0ms 在所有Go提交中击败了100.00%的用户
 * 占用内存：2MB 在所有Go提交中击败了59.35%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		x, _ := strconv.ParseFloat(scanner.Text(), 32)
		scanner.Scan()
		n, _ := strconv.Atoi(scanner.Text())
		Printf("Output: %v\n", myPow(x, n))
	}
}
