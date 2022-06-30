package main

import (
	"bufio"
	. "fmt"
	"os"
	"strconv"
	"strings"
)

// https://leetcode.cn/problems/license-key-formatting/
//------------------------Leetcode Problem 482------------------------
func licenseKeyFormatting(s string, k int) string {
	s = strings.ToUpper(strings.Replace(s, "-", "", -1)) //-1代表全部替换
	for i := len(s) - k; i > 0; i -= k {
		s = s[:i] + "-" + s[i:]
	}
	return s
}

//------------------------Leetcode Problem 482------------------------
/*
 * https://leetcode.cn/problems/license-key-formatting/
 * 执行用时：164ms 在所有Go提交中击败了48.51%的用户
 * 占用内存：6.8MB 在所有Go提交中击败了55.45%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		s := scanner.Text()
		scanner.Scan()
		k, _ := strconv.Atoi(scanner.Text())
		Printf("Output: %v\n", licenseKeyFormatting(s, k))
	}
}
