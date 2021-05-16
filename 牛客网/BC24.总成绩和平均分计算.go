package main

import "fmt"

func main() {
	var score1, score2, score3 float32
	fmt.Scanln(&score1, &score2, &score3)
	sum := score1 + score2 + score3
	fmt.Printf("%.2f %.2f", sum, sum/3)
}
/*
 * 运行时间：3ms 超过20.00%用Go提交的代码
 * 占用内存：804KB 超过100.00%用Go提交的代码
**/