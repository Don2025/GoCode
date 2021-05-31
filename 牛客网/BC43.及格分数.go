package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		n, _ := strconv.Atoi(input.Text())
		if n >= 60 {
			fmt.Println("Pass")
		} else {
			fmt.Println("Fail")
		}
	}
}

/*
 * 运行时间：4ms 超过62.50%用Go提交的代码
 * 占用内存：1004KB 超过12.50%用Go提交的代码
**/
