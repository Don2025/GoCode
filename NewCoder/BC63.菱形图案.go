package main

import "fmt"

func main() {
	for {
		var n int
		_, err := fmt.Scan(&n)
		if err != nil {
			return
		}
		n++
		for i := 1; i <= n; i++ {
			for j := 1; j <= n; j++ {
				if j > n-i {
					fmt.Print("* ")
				} else {
					fmt.Print(" ")
				}
			}
			fmt.Printf("\n")
		}
		for i := 2; i <= n; i++ {
			for j := 1; j <= n; j++ {
				if j < i {
					fmt.Print(" ")
				} else {
					fmt.Print("* ")
				}
			}
			fmt.Printf("\n")
		}
	}
}

/*
 * 运行时间：13ms 超过16.67%用Go提交的代码
 * 占用内存：880KB 超过0.00%用Go提交的代码
**/
