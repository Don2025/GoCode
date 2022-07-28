package main

import (
	"bufio"
	. "fmt"
	"os"
)

// https://leetcode.cn/problems/fraction-addition-and-subtraction/
func fractionAddition(expression string) string {

}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		Printf("Output: %v\n", fractionAddition(scanner.Text()))
	}
}
