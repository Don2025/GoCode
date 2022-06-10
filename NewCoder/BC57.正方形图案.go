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
			for j := 0; j < n; j++ {
				fmt.Print("* ")
			}
			fmt.Printf("\n")
		}
	}
}

/*
 * 运行时间：6ms 超过20.00%用Go提交的代码
 * 占用内存：956KB 超过20.00%用Go提交的代码
**/
