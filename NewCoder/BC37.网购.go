package main

import "fmt"

func main() {
	var price float64
	var month, date, flag int
	_, err := fmt.Scanln(&price, &month, &date, &flag)
	if err != nil {
		return
	}
	if month == 11 && date == 11 {
		price *= 0.7
	} else if month == 12 && date == 12 {
		price *= 0.8
	}
	if flag == 1 {
		price -= 50
	}
	if price < 0 {
		price = 0
	}
	fmt.Printf("%.2f", price)
}

/*
 * 运行时间：3ms 超过14.29%用Go提交的代码
 * 占用内存：912KB 超过42.86%用Go提交的代码
**/
