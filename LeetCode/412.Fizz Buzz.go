package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func fizzBuzz(n int) []string {
	var ans []string
	for i := 1; i <= n; i++ {
		if i%3 == 0 && i%5 == 0 {
			ans = append(ans, "FizzBuzz")
		} else if i%3 == 0 {
			ans = append(ans, "Fizz")
		} else if i%5 == 0 {
			ans = append(ans, "Buzz")
		} else {
			ans = append(ans, strconv.Itoa(i))
		}
	}
	return ans
}

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		n, _ := strconv.Atoi(input.Text())
		ans := fizzBuzz(n)
		for _, x := range ans {
			fmt.Printf("%s ", x)
		}
		fmt.Println()
	}
}

/*
 * 执行用时：4ms 在所有Go提交中击败了83.64%的用户
 * 占用内存：4.1MB 在所有Go提交中击败了29.37%的用户
**/
