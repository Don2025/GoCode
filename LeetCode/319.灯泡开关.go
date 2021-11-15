package main

import (
	"bufio"
	"math"
	"os"
	"strconv"
)

func bulbSwitch(n int) int {
	return int(math.Floor(math.Sqrt(float64(n))))
}

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		n, _ := strconv.Atoi(input.Text())
		println(bulbSwitch(n))
	}
}

/*
 * 执行用时：0ms 在所有Go提交中击败了100.00%的用户
 * 占用内存：1.9MB 在所有Go提交中击败了100.00%的用户
**/
