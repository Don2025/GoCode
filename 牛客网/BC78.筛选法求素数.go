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
		var cnt int
		for i := 2; i <= n; i++ {
			flag := true
			for j := 2; j < i; j++ {
				if i%j == 0 {
					flag = false
					cnt++
					break
				}
			}
			if flag {
				fmt.Printf("%d ", i)
			}
		}
		fmt.Printf("\n%d\n", cnt)
	}
}

/*
 * 运行时间：5ms 超过20.00%用Go提交的代码
 * 占用内存：992KB 超过20.00%用Go提交的代码
**/
