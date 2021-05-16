package main

import "fmt"

func main() {
	var n1, n2, n3, n4, n5 int
	_, err := fmt.Scanln(&n1, &n2, &n3, &n4, &n5)
	if err != nil {
		return
	}
	ave := float64(n1+n2+n3+n4+n5) / 5.0
	fmt.Printf("%.1f", ave)
}

/*
 * 运行时间：3ms 超过30.00%用Go提交的代码
 * 占用内存：880KB 超过60.00%用Go提交的代码
**/
