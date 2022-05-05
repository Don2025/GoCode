package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func add(a int, b int) int {
	var c int
	for b != 0 { // 进位不等于0
		c = (a & b) << 1 // 算出进位
		a ^= b           // 算出不带进位的和
		b = c            // 更新进位
	}
	return a
}

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		arr := strings.Fields(input.Text())
		a, _ := strconv.Atoi(arr[0])
		b, _ := strconv.Atoi(arr[1])
		println(add(a, b))
	}
}

/*
 * 执行用时：0ms 在所有Go提交中击败了100.00%的用户
 * 占用内存：1.8MB 在所有Go提交中击败了82.63%的用户
**/
