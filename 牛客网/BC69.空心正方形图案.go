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
				if i == 0 || j == 0 || i == n-1 || j == n-1 {
					fmt.Print("* ")
				} else {
					fmt.Print("  ")
				}
			}
			fmt.Printf("\n")
		}
	}
}

/*
 * 运行时间：6ms 超过16.67%用Go提交的代码
 * 占用内存：688KB 超过100.00%用Go提交的代码
**/
