package main

import (
	"fmt"
	"math"
)

func main() {
	for {
		var a, b, c float64
		_, err := fmt.Scanf("%f %f %f", &a, &b, &c)
		if err != nil {
			return
		}
		if a == 0 {
			fmt.Println("Not quadratic equation")
		} else {
			det := math.Pow(b, 2) - 4*a*c
			if det > 0 {
				fmt.Printf("x1=%.2f;x2=%.2f\n", (-b-math.Sqrt(b*b-4*a*c))/(2*a), (-b+math.Sqrt(b*b-4*a*c))/(2*a))
			} else if det < 0 {
				if b == 0 {
					fmt.Printf("x1=0.00-%.2fi;x2=0.00+%.2fi\n", math.Sqrt(-(b*b-4*a*c))/(2*a), math.Sqrt(-(b*b-4*a*c))/(2*a))
				} else {
					fmt.Printf("x1=%.2f-%.2fi;x2=%.2f+%.2fi\n", -b/(2*a), math.Sqrt(-(b*b-4*a*c))/(2*a), -b/(2*a), math.Sqrt(-(b*b-4*a*c))/(2*a))
				}
			} else {
				fmt.Printf("x1=x2=%.2f\n", -b/(2*a))
			}
		}
	}
}

/*
 * 运行时间：5ms 超过18.18%用Go提交的代码
 * 占用内存：972KB 超过9.09%用Go提交的代码
**/
