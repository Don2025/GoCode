package main

import "fmt"

func main() {
	var n1, n2 int
	for {
		_, err := fmt.Scanf("%d %d", &n1, &n2)
		if err != nil {
			return
		}
		if n1 > n2 {
			fmt.Printf("%d>%d\n", n1, n2)
		} else if n1 < n2 {
			fmt.Printf("%d<%d\n", n1, n2)
		} else {
			fmt.Printf("%d=%d\n", n1, n2)
		}
	}
}

/*
 * 运行时间：4ms 超过37.50%用Go提交的代码
 * 占用内存：888KB 超过25.00%用Go提交的代码
**/
