package main

import "fmt"

func main() {
	for i := 10000; i < 100000; i++ {
		sum := (i/10000)*(i%10000) + (i/1000)*(i%1000) + (i/100)*(i%100) + (i/10)*(i%10)
		if sum == i {
			fmt.Printf("%d ", i)
		}
	}
}

/*
 * 运行时间：3ms 超过50.00%用Go提交的代码
 * 占用内存：744KB 超过83.33%用Go提交的代码
**/
