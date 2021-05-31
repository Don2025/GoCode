package main

import "fmt"

func main() {
	var hour, minute, k int
	fmt.Scanf("%d:%d %d", &hour, &minute, &k)
	hour = (hour + (minute+k)/60) % 24
	minute = (minute + k) % 60
	fmt.Printf("%02d:%02d", hour, minute)
}

/*
 * 运行时间：2ms 超过100.00%用Go提交的代码
 * 占用内存：872KB 超过85.71%用Go提交的代码
**/
