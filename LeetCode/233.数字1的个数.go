package main

import (
	"bufio"
	"math"
	"os"
	"strconv"
)

func countDigitOne(n int) int {
	var cnt int
	j := 1
	//分别计算个十百...万位上1出现的次数再求和
	for i := n; i > 0; i /= 10 {
		if i%10 == 0 {
			cnt += (i / 10) * j
		}
		if i%10 == 1 {
			cnt += (i/10)*j + (n % j) + 1
		}
		if i%10 > 1 {
			cnt += int(math.Ceil(float64(i)/10)) * j
		}
		j *= 10
	}
	return cnt
}

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		n, _ := strconv.Atoi(input.Text())
		println(countDigitOne(n))
	}
}

/*
 * 执行用时：0ms 在所有Go提交中击败了100.00%的用户
 * 占用内存：1.9MB 在所有Go提交中击败了100.00%的用户
**/
