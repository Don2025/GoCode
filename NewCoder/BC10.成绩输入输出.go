package main

import "fmt"

func main() {
	var n1, n2, n3 int
	fmt.Scanln(&n1, &n2, &n3)
	fmt.Printf("score1=%d,score2=%d,score3=%d\n", n1, n2, n3)
}

/*
 * 运行时间：3ms 超过21.43%用Go提交的代码
 * 占用内存：920KB 超过28.57%用Go提交的代码
**/
