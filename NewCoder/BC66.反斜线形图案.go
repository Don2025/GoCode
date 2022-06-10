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
				if j == i {
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
 * 运行时间：6ms 超过60.00%用Go提交的代码
 * 占用内存：908KB 超过20.00%用Go提交的代码
**/
