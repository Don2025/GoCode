package main

import (
	"bufio"
	"math"
	"os"
	"strings"
)

func strToInt(str string) int {
	str = strings.TrimSpace(str)
	result := 0
	sign := 1
	for i, v := range str {
		if v >= '0' && v <= '9' {
			result = result*10 + int(v-'0')
		} else if v == '-' && i == 0 {
			sign = -1
		} else if v == '+' && i == 0 {
			sign = 1
		} else {
			break
		}
		if result > math.MaxInt32 {
			if sign == -1 {
				return math.MinInt32
			}
			return math.MaxInt32
		}
	}
	return sign * result
}

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		println(strToInt(input.Text()))
	}
}

/*
 * 执行用时：4ms 在所有Go提交中击败了40.17%的用户
 * 占用内存：2MB 在所有Go提交中击败了68.90%的用户
**/
