package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	input := bufio.NewScanner(os.Stdin)
	input.Scan()
	str := strings.ToLower(input.Text())
	input.Scan()
	ch := strings.ToLower(input.Text())
	var cnt int
	for _, val := range str {
		if val == rune(ch[0]) {
			cnt++
		}
	}
	fmt.Print(cnt)
}

/*
 * 运行时间：2ms 超过100.00%用Go提交的代码
 * 占用内存：964KB 超过40.24%用Go提交的代码
**/
