package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func selfDividingNumbers(left int, right int) []int {
	var isSelfDividing func(int) bool
	isSelfDividing = func(num int) bool {
		for x := num; x > 0; x /= 10 {
			if d := x % 10; d == 0 || num%d != 0 {
				return false
			}
		}
		return true
	}
	var ans []int
	for i := left; i <= right; i++ {
		if isSelfDividing(i) {
			ans = append(ans, i)
		}
	}
	return ans
}

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		left, _ := strconv.Atoi(input.Text())
		input.Scan()
		right, _ := strconv.Atoi(input.Text())
		ans := selfDividingNumbers(left, right)
		fmt.Printf("%v\n", ans)
	}
}

/*
 * 执行用时：0ms 在所有Go提交中击败了100.00%的用户
 * 占用内存：1.9MB 在所有Go提交中击败了59.09%的用户
**/
