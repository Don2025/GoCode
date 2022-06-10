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
		if unicode.IsLetter(rune(bytes[0])) {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}

/*
 * 运行时间：3ms 超过100.00%用Go提交的代码
 * 占用内存：960KB 超过62.50%用Go提交的代码
**/
