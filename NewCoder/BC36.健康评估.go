package main

import "fmt"

func main() {
	var weight, height float32
	fmt.Scanf("%f %f", &weight, &height)
	bmi := weight / (height * height)
	if bmi >= 18.5 && bmi <= 23.9 {
		fmt.Printf("Normal\n")
	} else {
		fmt.Printf("Abnormal\n")
	}
}

/*
 * 运行时间：3ms 超过14.29%用Go提交的代码
 * 占用内存：1048KB 超过14.29%用Go提交的代码
**/
