package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	input := bufio.NewReader(os.Stdin)
	text, _, _ := input.ReadLine()
	strArray := strings.Fields(string(text))
	if len(strArray) > 0 {
		fmt.Print(len(strArray[len(strArray)-1]))
	} else {
		fmt.Print(0)
	}
}

/*
 * 运行时间：2ms 超过85.31%用Go提交的代码
 * 占用内存：848KB 超过89.01%用Go提交的代码
**/
