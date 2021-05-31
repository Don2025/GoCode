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
			for j := 1; j <= n; j++ {
				if i == n-j+1 || i == j {
					fmt.Print("*")
				} else {
					fmt.Print(" ")
				}
			}
			fmt.Printf("\n")
		}
	}
}

/*
 * 运行时间：8ms 超过0.00%用Go提交的代码
 * 占用内存：996KB 超过0.00%用Go提交的代码
**/
