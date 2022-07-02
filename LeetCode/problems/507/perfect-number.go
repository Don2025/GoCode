package main

import (
	"bufio"
	. "fmt"
	"os"
	"strconv"
)

// https://leetcode.cn/problems/perfect-number/
//------------------------Leetcode Problem 507------------------------
func checkPerfectNumber(num int) bool {
	if num == 1 {
		return false
	}
	sum := 1
	for i := 2; i*i <= num; i++ {
		if num%i == 0 {
			sum += i
			if i*i < num {
				sum += num / i
			}
		}
	}
	return sum == num
}

//------------------------Leetcode Problem 507------------------------
/*
 * https://leetcode.cn/problems/perfect-number/
 * 执行用时：100ms 在所有Go提交中击败了100.00%的用户
 * 占用内存：1.9MB 在所有Go提交中击败了100.00%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		num, _ := strconv.Atoi(scanner.Text())
		Printf("Output: %v\n", checkPerfectNumber(num))
	}
}
