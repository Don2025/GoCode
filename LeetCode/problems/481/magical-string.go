package main

import (
	"bufio"
	"bytes"
	. "fmt"
	"os"
	"strconv"
)

// https://leetcode.cn/problems/magical-string/
//------------------------Leetcode Problem 481------------------------
func magicalString(n int) int {
	s := make([]byte, 0, n+1)
	s = append(s, 1, 2, 2)
	c := []byte{2}
	for i := 2; len(s) < n; i++ {
		c[0] ^= 3
		s = append(s, bytes.Repeat(c, int(s[i]))...)
	}
	return bytes.Count(s[:n], []byte{1})
}

//------------------------Leetcode Problem 481------------------------
/*
 * https://leetcode.cn/problems/magical-string/
 * 执行用时：8ms 在所有Go提交中击败了30.77%的用户
 * 占用内存：2.9MB 在所有Go提交中击败了76.92%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		n, _ := strconv.Atoi(scanner.Text())
		Printf("Output: %v", magicalString(n))
	}
}
