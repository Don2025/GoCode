package main

import (
	"bufio"
	. "fmt"
	"os"
	"strconv"
)

// https://leetcode.cn/problems/subtract-the-product-and-sum-of-digits-of-an-integer/
//------------------------Leetcode Problem 1281------------------------
func subtractProductAndSum(n int) int {
	add, mul := 0, 1
	for i := n; i > 0; i /= 10 {
		t := i % 10
		add += t
		mul *= t
	}
	return mul - add
}

//------------------------Leetcode Problem 1281------------------------
/*
 * https://leetcode.cn/problems/subtract-the-product-and-sum-of-digits-of-an-integer/
 * 执行用时：0ms 在所有Go提交中击败了100.00%的用户
 * 占用内存：1.8MB 在所有Go提交中击败了58.95%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		n, _ := strconv.Atoi(scanner.Text())
		Printf("Output: %v\n", subtractProductAndSum(n))
	}
}
