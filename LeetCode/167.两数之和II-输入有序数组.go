package main

import (
	"bufio"
	. "fmt"
	"github.com/Don2025/GoCode/utils"
	"os"
	"strconv"
)

func twoSum(numbers []int, target int) []int {
	l, r := 0, len(numbers)-1
	for l < r {
		sum := numbers[l] + numbers[r]
		if target == sum {
			return []int{l + 1, r + 1}
		} else if target < sum {
			r--
		} else { // target > sum
			l++
		}
	}
	return []int{}
}

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		nums := utils.StringToInts(input.Text())
		input.Scan()
		target, _ := strconv.Atoi(input.Text())
		Printf("Output: %v\n", twoSum(nums, target))
	}
}

/*
 * 执行用时：4ms 在所有Go提交中击败了86.97%的用户
 * 占用内存：3MB 在所有Go提交中击败了67.99%的用户
**/
