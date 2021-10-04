package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func licenseKeyFormatting(s string, k int) string {
	s = strings.ToUpper(strings.Replace(s, "-", "", -1)) //-1代表全部替换
	for i := len(s) - k; i > 0; i -= k {
		s = s[:i] + "-" + s[i:]
	}
	return s
}

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		s := input.Text()
		input.Scan()
		k, _ := strconv.Atoi(input.Text())
		println(licenseKeyFormatting(s, k))
	}
}

/*
 * 执行用时：164ms 在所有Go提交中击败了48.51%的用户
 * 占用内存：6.8MB 在所有Go提交中击败了55.45%的用户
**/
