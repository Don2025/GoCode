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
		fmt.Println(1 << n)
	}
}

/*
 * 运行时间：3ms 超过27.27%用Go提交的代码
 * 占用内存：920KB 超过45.45%用Go提交的代码
**/
