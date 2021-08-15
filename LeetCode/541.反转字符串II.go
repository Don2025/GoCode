package main

import (
	"bufio"
	"os"
	"strconv"
)

func reverseStr(s string, k int) string {
	arr := []byte(s)
	for i := 0; i < len(arr); i += 2 * k {
		r := i + k - 1
		if r >= len(arr) {
			r = len(arr) - 1
		}
		mid := (i + r) >> 1
		for j := 0; j <= mid-i; j++ {
			arr[i+j], arr[r-j] = arr[r-j], arr[i+j]
		}
	}
	return string(arr)
}

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		s := input.Text()
		input.Scan()
		k, _ := strconv.Atoi(input.Text())
		println(reverseStr(s, k))
	}
}

/*
 * 执行用时：0ms 在所有Go提交中击败了100.00%的用户
 * 占用内存：3.4MB 在所有Go提交中击败了100.00%的用户
**/
