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
		switch n {
		case 200:
			fmt.Println("OK")
		case 202:
			fmt.Println("Accepted")
		case 400:
			fmt.Println("Bad Request")
		case 403:
			fmt.Println("Forbidden")
		case 404:
			fmt.Println("Not Found")
		case 500:
			fmt.Println("Internal Server Error")
		case 502:
			fmt.Println("Bad Gateway")
		default:

		}
	}
}

/*
 * 运行时间：3ms 超过12.50%用Go提交的代码
 * 占用内存：936KB 超过12.50%用Go提交的代码
**/
