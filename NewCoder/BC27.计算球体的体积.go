package main

import (
	"fmt"
	"math"
)

const PI = 3.1415926

func main() {
	var r float64
	_, err := fmt.Scan(&r)
	if err != nil {
		return
	}
	v := 4 / 3.0 * PI * math.Pow(r, 3)
	fmt.Printf("%.3f", v)
}

/*
 * 运行时间：3ms 超过8.33%用Go提交的代码
 * 占用内存：868KB 超过50.00%用Go提交的代码
**/
