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
		if n >= 140 {
			fmt.Println("Genius")
		}
	}
}

/*
 * 运行时间：3ms 超过14.29%用Go提交的代码
 * 占用内存：1064KB 超过0.00%用Go提交的代码
**/
