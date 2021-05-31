package main

import "fmt"

func main() {
	for {
		var n int
		_, err := fmt.Scan(&n)
		if err != nil {
			return
		}
		for i := 0; i < n; i++ {
			fmt.Print("*")
		}
		fmt.Printf("\n")
	}
}

/*
 * 运行时间：3ms 超过20.00%用Go提交的代码
 * 占用内存：828KB 超过40.00%用Go提交的代码
**/
