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
 * 运行时间：11ms 超过9.09%用Go提交的代码
 * 占用内存：972KB 超过18.18%用Go提交的代码
**/
