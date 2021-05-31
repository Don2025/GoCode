package main

import "fmt"

func main() {
	var a, b, sum float64
	var c rune
	for {
		_, err := fmt.Scanf("%f%c%f", &a, &c, &b)
		if err != nil {
			return
		}
		switch c {
		case '+':
			sum = a + b
			fmt.Printf("%.4f+%.4f=%.4f\n", a, b, sum)
		case '-':
			sum = a - b
			fmt.Printf("%.4f-%.4f=%.4f\n", a, b, sum)
		case '*':
			sum = a * b
			fmt.Printf("%.4f*%.4f=%.4f\n", a, b, sum)
		case '/':
			sum = a / b
			if b == 0 {
				fmt.Println("Wrong!Division by zero!")
			} else {
				fmt.Printf("%.4f/%.4f=%.4f\n", a, b, sum)
			}
		default:
			fmt.Println("Invalid operation!")
		}
	}
}

/*
 * 运行时间：4ms 超过28.57%用Go提交的代码
 * 占用内存：980KB 超过14.29%用Go提交的代码
**/
