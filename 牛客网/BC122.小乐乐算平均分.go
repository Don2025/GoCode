package main

import "fmt"

func main() {
	var a, b, c float64
	fmt.Scan(&a, &b, &c)
	sum := a + b + c
	fmt.Printf("%.2f %.2f\n", sum, sum/3)
}

/*
 * 运行时间：4ms 超过12.50%用Go提交的代码
 * 占用内存：908KB 超过37.50%用Go提交的代码
**/
