package main

import (
	"bufio"
	"os"
	"strconv"
)

/**
 *
 * @param x int整型
 * @return bool布尔型
 */
func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	str := strconv.Itoa(x)
	for i := 0; i < len(str)/2; i++ {
		if str[i] != str[len(str)-1-i] {
			return false
		}
	}
	return true
}

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		n, _ := strconv.Atoi(input.Text())
		println(isPalindrome(n))
	}
}

/*
 * 运行时间：3ms 超过37.65%用Go提交的代码
 * 占用内存：1168KB 超过9.57%用Go提交的代码
**/
