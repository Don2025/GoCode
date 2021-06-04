package main

import (
	"bufio"
	"fmt"
	"os"
)

//在小写字母字符串 s 中找出第一个只出现一次的字符。若没有则返回单空格。
func firstUniqChar(s string) byte {
	hash := [26]int{}
	for _, x := range s {
		hash[x-'a']++
	}
	for _, x := range s {
		if hash[x-'a'] == 1 {
			return byte(x)
		}
	}
	return ' '
}

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		fmt.Printf("%c\n", firstUniqChar(input.Text()))
	}
}

/*
 * 执行用时：4ms 在所有Go提交中击败了97.55%的用户
 * 占用内存：5.3MB 在所有Go提交中击败了76.72%的用户
**/
