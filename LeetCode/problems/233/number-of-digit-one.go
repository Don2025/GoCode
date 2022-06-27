package main

import (
	"bufio"
	. "fmt"
	"math"
	"os"
	"strconv"
)

// https://leetcode.cn/problems/number-of-digit-one/
//------------------------Leetcode Problem 233------------------------
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

//------------------------Leetcode Problem 233------------------------
/*
 * https://leetcode.cn/problems/number-of-digit-one/
 * 执行用时：0ms 在所有Go提交中击败了100.00%的用户
 * 占用内存：1.9MB 在所有Go提交中击败了100.00%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		n, _ := strconv.Atoi(scanner.Text())
		Printf("Output: %v\n", countDigitOne(n))
	}
}
