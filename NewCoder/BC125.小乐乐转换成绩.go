package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	if n < 0 || n > 100 {
		fmt.Print("F")
	} else if n < 60 {
		fmt.Print("E")
	} else if n < 70 {
		fmt.Print("D")
	} else if n < 80 {
		fmt.Print("C")
	} else if n < 90 {
		fmt.Print("B")
	} else {
		fmt.Print("A")
	}
}

/*
 * 运行时间：5ms 超过0.00%用Go提交的代码
 * 占用内存：908KB 超过50.00%用Go提交的代码
**/
