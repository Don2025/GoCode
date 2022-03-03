package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func addToArrayForm(num []int, k int) []int {
	carry, i := 0, len(num)-1
	for k+carry != 0 {
		if i >= 0 {
			num[i] += k%10 + carry
			num[i], carry, k, i = num[i]%10, num[i]/10, k/10, i-1
		} else {
			num = append([]int{carry + k%10}, num...)
			num[0], carry, k = num[0]%10, num[0]/10, k/10
		}
	}
	return num
}

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		num := stringArrayToIntArray(strings.Fields(input.Text()))
		input.Scan()
		k, _ := strconv.Atoi(input.Text())
		println(addToArrayForm(num, k))
	}
}

/*
 * 执行用时：20ms 在所有Go提交中击败了92.55%的用户
 * 占用内存：6.4MB 在所有Go提交中击败了93.62%的用户
**/
