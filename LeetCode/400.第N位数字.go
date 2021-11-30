package main

import (
	"bufio"
	"math"
	"os"
	"strconv"
)

func findNthDigit(n int) int {
	d := 1
	for cnt := 9; n > d*cnt; cnt *= 10 {
		n -= d * cnt
		d++
	}
	index := n - 1
	start := int(math.Pow10(d - 1))
	num := start + index/d
	digitIndex := index % d
	return num / int(math.Pow10(d-digitIndex-1)) % 10
}

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		n, _ := strconv.Atoi(input.Text())
		println(findNthDigit(n))
	}
}

/*
 * 执行用时：0ms 在所有Go提交中击败了100.00%的用户
 * 占用内存：1.9MB 在所有Go提交中击败了37.33%的用户
**/
