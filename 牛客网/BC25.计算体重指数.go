package main

import (
	"fmt"
	"math"
)

func main() {
	var height, weight float64
	fmt.Scanf("%f %f", &height, &weight)
	fmt.Printf("%.2f", height/math.Pow(weight/100,2))
}
/*
 * 运行时间：3ms 超过16.67%用Go提交的代码
 * 占用内存：900KB 超过33.33%用Go提交的代码
**/