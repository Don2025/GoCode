package main

import (
	"fmt"
	"math"
)

func main() {
	var age int
	fmt.Scan(&age)
	t := 3.156 * (math.Pow(10, 7))
	fmt.Println(age * int(t))
}

/*
 * 运行时间：3ms 超过14.29%用Go提交的代码
 * 占用内存：868KB 超过42.86%用Go提交的代码
**/
