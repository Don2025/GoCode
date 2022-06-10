package main

import "fmt"

func main() {
	var y, m, d int
	fmt.Scan(&y, &m, &d)
	fmt.Printf("%d/%d/%d\n", d, m, y)
}

/*
 * 运行时间：3ms 超过18.18%用Go提交的代码
 * 占用内存：952KB 超过54.55%用Go提交的代码
**/
