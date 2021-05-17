package main

import "fmt"

func main() {
	for {
		var weight, height float64
		_, err := fmt.Scanln(&weight, &height)
		if err != nil {
			return
		}
		height /= 100
		weight /= height * height
		if weight < 18.5 {
			fmt.Println("Underweight")
		} else if weight <= 23.9 {
			fmt.Println("Normal")
		} else if weight <= 27.9 {
			fmt.Println("Overweight")
		} else {
			fmt.Println("Obese")
		}
	}
}

/*
 * 运行时间：5ms 超过33.33%用Go提交的代码
 * 占用内存：1012KB 超过0.00%用Go提交的代码
**/
