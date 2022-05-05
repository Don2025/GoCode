package main

import (
	"bufio"
	"math"
	"os"
	"strconv"
)

func cuttingRope(n int) int {
	if n <= 3 {
		return n - 1
	}
	return int(math.Pow(3, float64(n/3)) * 4.0 / float64(4-n%3))
}

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		n, _ := strconv.Atoi(input.Text())
		println(cuttingRope(n))
	}
}

/*
 * 执行用时：0ms 在所有Go提交中击败了100.00%的用户
 * 占用内存：1.9MB 在所有Go提交中击败了41.73%的用户
**/
