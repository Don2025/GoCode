package main

import "fmt"

func main() {
	for {
		var n int
		_, err := fmt.Scan(&n)
		if err != nil {
			return
		}
		for i := 1; i <= n; i++ {
			for j := 1; j <= i; j++ {
				fmt.Print("* ")
			}
			fmt.Printf("\n")
		}
	}
}

/*
 * 运行时间：5ms 超过0.00%用Go提交的代码
 * 占用内存：940KB 超过16.67%用Go提交的代码
**/
