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
		for i := 1; i <= n; i++ {
			for j := 1; j <= i; j++ {
				fmt.Printf("%d ", j)
			}
			fmt.Printf("\n")
		}
	}
}

/*
 * 运行时间：5ms 超过100.00%用Go提交的代码
 * 占用内存：876KB 超过28.57%用Go提交的代码
**/
