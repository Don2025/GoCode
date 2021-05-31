package main

import "fmt"

func main() {
	var n int
	_, err := fmt.Scan(&n)
	if err != nil {
		return
	}
	if n >= 10 {
		fmt.Print("Danger++")
	} else if n >= 4 {
		fmt.Print("Danger")
	} else {
		fmt.Print("Good")
	}
}

/*
 * 运行时间：7ms 超过0.00%用Go提交的代码
 * 占用内存：888KB 超过100.00%用Go提交的代码
**/
