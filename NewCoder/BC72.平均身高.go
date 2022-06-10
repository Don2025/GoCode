package main

import "fmt"

func main() {
	var h1, h2, h3, h4, h5 float64
	_, err := fmt.Scanf("%f %f %f %f %f", &h1, &h2, &h3, &h4, &h5)
	if err != nil {
		return
	}
	ave := (h1 + h2 + h3 + h4 + h5) / 5
	fmt.Printf("%.2f", ave)
}

/*
 * 运行时间：8ms 超过25.00%用Go提交的代码
 * 占用内存：928KB 超过75.00%用Go提交的代码
**/
