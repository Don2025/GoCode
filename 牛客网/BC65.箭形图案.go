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
			for j := 0; j < n-i; j++ {
				fmt.Print("  ")
			}
			for j := 0; j <= i; j++ {
				fmt.Print("*")
			}
			fmt.Printf("\n")
		}
		for i := 0; i < n+1; i++ {
			fmt.Print("*")
		}
		fmt.Printf("\n")
		for i := 0; i < n; i++ {
			for j := 0; j <= i; j++ {
				fmt.Print("  ")
			}
			for j := 0; j < n-i; j++ {
				fmt.Print("*")
			}
			fmt.Printf("\n")
		}
	}
}

/*
 * 运行时间：11ms 超过37.50%用Go提交的代码
 * 占用内存：888KB 超过25.00%用Go提交的代码
**/
