package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		strArray := strings.Fields(input.Text())
		for i := len(strArray) - 1; i >= 0; i-- {
			fmt.Printf("%s ", strArray[i])
		}
	}
}

/*
 * 运行时间：3ms 超过56.98%用Go提交的代码
 * 占用内存：1020KB 超过8.92%用Go提交的代码
**/
