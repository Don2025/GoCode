package main

import (
	"bufio"
	"github.com/Don2025/GoCode/utils"
	"os"
)

func singleNumber(nums []int) int {
	var ans int
	for _, x := range nums {
		ans ^= x
	}
	return ans
}

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		println(singleNumber(utils.StringToInts(input.Text())))
	}
}

/*
 * 执行用时：12ms 在所有Go提交中击败了93.66%的用户
 * 占用内存：6MB 在所有Go提交中击败了78.66%的用户
**/
