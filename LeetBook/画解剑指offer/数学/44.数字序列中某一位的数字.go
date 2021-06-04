package main

import (
	"bufio"
	"os"
	"strconv"
)

//数字以0123456789101112131415…的格式序列化到一个字符序列中。请写一个函数，求任意第n位对应的数字。
func findNthDigit(n int) int {
	digit, start, cnt := 1, 1, 9
	for n-cnt > 0 {
		n -= cnt
		digit++
		start *= 10
		cnt = digit * start * 9
	}
	num := start + (n-1)/digit
	return int(strconv.Itoa(num)[(n-1)%digit] - '0')
}

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		n, _ := strconv.Atoi(input.Text())
		println(findNthDigit(n))
	}
}

/*
 * 执行用时：0ms 在所有Go提交中击败了100.00%的用户
 * 占用内存：1.9MB 在所有Go提交中击败了50.66%的用户
**/
