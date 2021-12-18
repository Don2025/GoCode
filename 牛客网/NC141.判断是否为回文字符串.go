package main

import (
	"bufio"
	"os"
)

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 * @param str string字符串 待判断的字符串
 * @return bool布尔型
 */
func judge(str string) bool {
	for i := 0; i < len(str)/2; i++ {
		if str[i] != str[len(str)-1-i] {
			return false
		}
	}
	return true
}

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		println(judge(input.Text()))
	}
}

/*
 * 运行时间：10ms 超过13.74%用Go提交的代码
 * 占用内存：3344KB 超过95.24%用Go提交的代码
**/
