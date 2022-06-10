package main

import (
	"bufio"
	"os"
)

//请实现一个函数，把字符串 s 中的每个空格替换成"%20"。
func replaceSpace(s string) string {
	str := make([]rune, 0)
	for _, it := range s {
		if it == ' ' {
			str = append(str, '%', '2', '0')
		} else {
			str = append(str, it)
		}
	}
	return string(str)
}

/* 一种不要脸的简洁写法
func replaceSpace(s string) string {
	return strings.ReplaceAll(s, " ", "%20")
}
*/

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		println(replaceSpace(input.Text()))
	}
}

/*
 * 执行用时：0ms 在所有Go提交中击败了100.00%的用户
 * 占用内存：2MB 在所有Go提交中击败了30.44%的用户
**/
