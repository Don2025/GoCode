package main

import (
	"bufio"
	"os"
	"regexp"
)

func isNumber(s string) bool {
	ans, _ := regexp.MatchString(`^\s*[+-]?((\d*\.\d+)|(\d+\.?\d*))([eE][+-]?\d+)?\s*$`, s)
	return ans
}

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		println(isNumber(input.Text()))
	}
}

/*
 * 执行用时：20ms 在所有Go提交中击败了6.35%的用户
 * 占用内存：7MB 在所有Go提交中击败了5.69%的用户
**/
