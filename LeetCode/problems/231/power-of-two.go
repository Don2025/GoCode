package main

import (
	"bufio"
	. "fmt"
	"os"
	"strconv"
)

// https://leetcode.cn/problems/power-of-two/
//------------------------Leetcode Problem 231------------------------
func isPowerOfTwo(n int) bool {
	//利用负数在计算机中的存储方式，检测一个数是否为2的幂只要判断 n&(-n) 是否等于 n 本身即可
	return n > 0 && (n&(-n)) == n
}

//------------------------Leetcode Problem 231------------------------
/*
 * https://leetcode.cn/problems/power-of-two/
 * 执行用时：0ms 在所有Go提交中击败了100.00%的用户
 * 占用内存：2.1MB 在所有Go提交中击败了31.95%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		n, _ := strconv.Atoi(scanner.Text())
		Printf("Output: %v\n", isPowerOfTwo(n))
	}
}
