package main

import "fmt"

func main() {
	for {
		var year, month int
		_, err := fmt.Scanf("%d %d", &year, &month)
		if err != nil {
			return
		}
		date := []int{0, 31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
		if leapYear(year) && month == 2 {
			fmt.Printf("%d\n", date[month]+1)
		} else {
			fmt.Printf("%d\n", date[month])
		}
	}
}

func leapYear(year int) bool {
	return year%400 == 0 || (year%100 != 0 && year%4 == 0)
}

/*
 * 运行时间：4ms 超过20.00%用Go提交的代码
 * 占用内存：964KB 超过20.00%用Go提交的代码
**/
