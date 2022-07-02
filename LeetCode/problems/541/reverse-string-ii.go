package main

import (
	"bufio"
	. "fmt"
	"os"
	"strconv"
)

// https://leetcode.cn/problems/reverse-string-ii/
//------------------------Leetcode Problem 541------------------------
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

//------------------------Leetcode Problem 541------------------------
/*
 * https://leetcode.cn/problems/reverse-string-ii/
 * 执行用时：0ms 在所有Go提交中击败了100.00%的用户
 * 占用内存：3.4MB 在所有Go提交中击败了100.00%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		s := scanner.Text()
		scanner.Scan()
		k, _ := strconv.Atoi(scanner.Text())
		Printf("Output: %v\n", reverseStr(s, k))
	}
}
