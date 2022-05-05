package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func addStrings(num1 string, num2 string) string {
	var add int
	var ans string
	for i, j := len(num1)-1, len(num2)-1; i >= 0 || j >= 0 || add != 0; i, j = i-1, j-1 {
		var x, y int
		if i >= 0 {
			x = int(num1[i] - '0')
		}
		if j >= 0 {
			y = int(num2[j] - '0')
		}
		n := x + y + add
		ans = strconv.Itoa(n%10) + ans
		add = n / 10
	}
	return ans
}

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		arr := strings.Fields(input.Text())
		println(addStrings(arr[0], arr[1]))
	}
}

/*
 * 执行用时：4ms 在所有Go提交中击败了77.30%的用户
 * 占用内存：6.8MB 在所有Go提交中击败了58.89%的用户
**/
