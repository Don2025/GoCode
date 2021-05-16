package main

import "fmt"

func main() {
	var(
		stuId int
		score1 float32
		score2 float32
		score3 float32
	)
	fmt.Scanf("%d;%f,%f,%f", &stuId, &score1, &score2, &score3)
	fmt.Printf("The each subject score of  No. %d is %.2f, %.2f, %.2f.\n", stuId, score1, score2, score3)
}
/*
 * 运行时间：3ms 超过20.69%用Go提交的代码
 * 占用内存：828KB 超过82.76%用Go提交的代码
**/