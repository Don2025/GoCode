package main

import (
	"bufio"
	. "fmt"
	"os"
	"strconv"
	"strings"
)

// https://leetcode.cn/problems/generate-a-string-with-characters-that-have-odd-counts/
//------------------------Leetcode Problem 1374------------------------
func generateTheString(n int) string {
	if n&1 == 0 {
		return strings.Repeat("s", n-1) + "b"
	} else {
		return strings.Repeat("s", n)
	}
}

//------------------------Leetcode Problem 1374------------------------
/*
 * https://leetcode.cn/problems/generate-a-string-with-characters-that-have-odd-counts/
 * 执行用时：0ms 在所有Go提交中击败了100.00%的用户
 * 占用内存：2.1MB 在所有Go提交中击败了81.25%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		n, _ := strconv.Atoi(scanner.Text())
		Printf("Output: %v\n", generateTheString(n))
	}
}
