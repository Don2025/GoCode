package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		fmt.Printf("%d\n", xToDec(string(input.Text())[2:], 16))
	}
}

//将给定的x进制字符串s转换成10进制
func xToDec(s string, x int) (ans int) {
	for i := 0; i < len(s); i++ {
		if unicode.IsNumber(rune(s[i])) {
			ans = ans*x + int(s[i]-'0')
		} else {
			ans = ans*x + int(s[i]-'A'+10)
		}
	}
	return
}

/*
 * 运行时间：2ms 超过100.00%用Go提交的代码
 * 占用内存：852KB 超过91.51%用Go提交的代码
**/
