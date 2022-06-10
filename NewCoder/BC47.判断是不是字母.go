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
		if unicode.IsLetter(ch) {
			fmt.Printf("%c is an alphabet.\n", ch)
		} else {
			fmt.Printf("%c is not an alphabet.\n", ch)
		}
	}
}

/*
 * 运行时间：9ms 超过0.00%用Go提交的代码
 * 占用内存：964KB 超过0.00%用Go提交的代码
**/
