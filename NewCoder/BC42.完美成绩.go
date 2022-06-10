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
		if n >= 90 && n <= 100 {
			fmt.Println("Perfect")
		}
	}
}

/*
 * 运行时间：5ms 超过10.00%用Go提交的代码
 * 占用内存：892KB 超过80.00%用Go提交的代码
**/
