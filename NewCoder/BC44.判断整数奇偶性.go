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
		if n%2 == 0 {
			fmt.Println("Even")
		} else {
			fmt.Println("Odd")
		}
	}
}

/*
 * 运行时间：3ms 超过50.00%用Go提交的代码
 * 占用内存：900KB 超过50.00%用Go提交的代码
**/
