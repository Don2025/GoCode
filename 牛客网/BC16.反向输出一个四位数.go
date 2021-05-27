package main

import "fmt"

func main() {
	var n, sum int
	fmt.Scan(&n)
	for n != 0 {
		sum = sum*10 + n%10
		n /= 10
	}
	fmt.Printf("%04d\n", sum)
}
