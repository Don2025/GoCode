package main

import (
	"bufio"
	"os"
)

func titleToNumber(columnTitle string) int {
	var ans int
	for _, x := range columnTitle {
		ans *= 26
		ans += int(x - 'A' + 1)
	}
	return ans
}

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		println(titleToNumber(input.Text()))
	}
}
