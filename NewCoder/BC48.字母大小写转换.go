package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		bytes := input.Bytes()
		var ch rune = rune(bytes[0])
		if unicode.IsLower(ch) {
			fmt.Printf("%c\n", ch-32)
		} else {
			fmt.Printf("%c\n", ch+32)
		}
	}
}

/*
 * 运行时间：4ms 超过0.00%用Go提交的代码
 * 占用内存：896KB 超过66.67%用Go提交的代码
**/
