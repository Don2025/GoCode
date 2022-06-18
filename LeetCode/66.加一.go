package main

import (
	"bufio"
	. "fmt"
	"github.com/Don2025/GoCode/utils"
	"os"
)

func plusOne(digits []int) []int {
	for i := len(digits) - 1; i >= 0; i-- {
		if digits[i] != 9 {
			digits[i]++
			return digits
		}
		digits[i] = 0
	}
	digits = make([]int, len(digits)+1)
	digits[0] = 1
	return digits
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		digit := utils.StringToInts(scanner.Text())
		Printf("Output: %v\n", plusOne(digit))
	}
}

/*
 * 执行用时：0ms 在所有Go提交中击败了100.00%的用户
 * 占用内存：2MB 在所有Go提交中击败了56.24%的用户
**/
