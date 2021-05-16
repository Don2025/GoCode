package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b, c, cir, area float64
	fmt.Scanln(&a, &b, &c)
	cir = a + b + c
	q := cir / 2
	area = math.Sqrt(q*(q-a)*(q-b)*(q-c))
	fmt.Printf("circumference=%.2f area=%.2f", a+b+c, area)
}
/*
 * 运行时间：3ms 超过18.75%用Go提交的代码
 * 占用内存：816KB 超过87.50%用Go提交的代码
**/