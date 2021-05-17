package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		text := input.Text()
		switch text {
		case "A", "E", "I", "O", "U", "a", "e", "i", "o", "u":
			fmt.Println("Vowel")
		default:
			fmt.Println("Consonant")
		}
	}
}

/*
 * 运行时间：4ms 超过0.00%用Go提交的代码
 * 占用内存：1016KB 超过0.00%用Go提交的代码
**/
