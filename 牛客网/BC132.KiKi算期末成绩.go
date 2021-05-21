package main

import "fmt"

func main() {
	var a, b, c, d float64
	fmt.Scan(&a, &b, &c, &d)
	fmt.Printf("%.1f", a*0.2+b*0.1+c*0.2+d*0.5)
}

/*
 * 运行时间：2ms 超过100.00%用Go提交的代码
 * 占用内存：888KB 超过62.50%用Go提交的代码
**/
