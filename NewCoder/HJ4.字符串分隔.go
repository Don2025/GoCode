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
		for len(text) > 8 {
			fmt.Println(text[:8])
			text = text[8:]
		}
		if len(text) <= 8 {
			for i := len(text); i < 8; i++ {
				text += "0"
			}
			fmt.Println(text)
			continue
		}
	}
}

/*
 * 运行时间：2ms 超过100.00%用Go提交的代码
 * 占用内存：892KB 超过59.28%用Go提交的代码
**/
