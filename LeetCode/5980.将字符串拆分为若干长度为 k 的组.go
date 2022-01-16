package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func divideString(s string, k int, fill byte) []string {
	if t := len(s) % k; t != 0 {
		for i := k - t; i > 0; i-- {
			s += string(fill)
		}
	}
	var ans []string
	for i := 0; i < len(s); i += k {
		ans = append(ans, s[i:i+k])
	}
	return ans
}

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		s := input.Text()
		input.Scan()
		k, _ := strconv.Atoi(input.Text())
		input.Scan()
		fill := byte(input.Text()[0])
		ans := divideString(s, k, fill)
		for _, x := range ans {
			fmt.Printf("%s ", x)
		}
		fmt.Println()
	}
}
