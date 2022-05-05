package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func rangeBitwiseAnd(left int, right int) int {
	var shift int
	for left < right {
		left >>= 1
		right >>= 1
		shift++
	}
	return left << shift
}

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		arr := strings.Fields(input.Text())
		left, _ := strconv.Atoi(arr[0])
		right, _ := strconv.Atoi(arr[1])
		println(rangeBitwiseAnd(left, right))
	}
}

/*
 * 执行用时：4ms 在所有Go提交中击败了97.94%的用户
 * 占用内存：5MB 在所有Go提交中击败了63.79%的用户
**/
