package main

import "fmt"

func main() {
	for {
		var n int
		_, err := fmt.Scanf("%d", &n)
		if err != nil {
			return
		}
		if n > 0 {
			fmt.Println(1)
		} else if n < 0 {
			fmt.Println(0)
		} else {
			fmt.Println(0.5)
		}
	}
}

/*
 * 运行时间：3ms 超过33.33%用Go提交的代码
 * 占用内存：888KB 超过83.33%用Go提交的代码
**/
