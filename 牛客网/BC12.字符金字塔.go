package main

import "fmt"

func main() {
	var ch string
	fmt.Scan(&ch)
	for i := 0; i < 5; i++ {
		for j := 0; j < 4-i; j++ {
			fmt.Printf(" ")
		}
		for j := 0; j < i+1; j++ {
			fmt.Printf("%s ", ch)
		}
		fmt.Printf("\n")
	}
}

/*
 * 运行时间：6ms 超过15.79%用Go提交的代码
 * 占用内存：784KB 超过47.37%用Go提交的代码
**/
