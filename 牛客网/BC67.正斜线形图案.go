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
				if i == n-j+1 {
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
 * 运行时间：7ms 超过0.00%用Go提交的代码
 * 占用内存：768KB 超过40.00%用Go提交的代码
**/
