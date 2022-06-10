package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		bytes := input.Bytes()
		fmt.Printf("%c\n", bytes[0]+32)
	}
}

/*
 * 运行时间：2ms 超过100.00%用Go提交的代码
 * 占用内存：892KB 超过52.94%用Go提交的代码
**/
