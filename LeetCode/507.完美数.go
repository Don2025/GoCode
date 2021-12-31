package main

import (
	"bufio"
	"os"
	"strconv"
)

func checkPerfectNumber(num int) bool {
	if num == 1 {
		return false
	}
	sum := 1
	for i := 2; i*i <= num; i++ {
		if num%i == 0 {
			sum += i
			if i*i < num {
				sum += num / i
			}
		}
	}
	return sum == num
}

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		num, _ := strconv.Atoi(input.Text())
		println(checkPerfectNumber(num))
	}
}

/*
 * 执行用时：100ms 在所有Go提交中击败了100.00%的用户
 * 占用内存：1.9MB 在所有Go提交中击败了100.00%的用户
**/
