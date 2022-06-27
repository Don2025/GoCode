package main

import (
	"bufio"
	. "fmt"
	"os"
	"strconv"
)

// https://leetcode.cn/problems/fraction-to-recurring-decimal/
//------------------------Leetcode Problem 166------------------------
func fractionToDecimal(numerator, denominator int) string {
	if numerator%denominator == 0 {
		return strconv.Itoa(numerator / denominator)
	}

	s := []byte{}
	if numerator < 0 != (denominator < 0) {
		s = append(s, '-')
	}
	// 整数部分
	numerator = abs(numerator)
	denominator = abs(denominator)
	integerPart := numerator / denominator
	s = append(s, strconv.Itoa(integerPart)...)
	s = append(s, '.')
	// 小数部分
	indexMap := map[int]int{}
	remainder := numerator % denominator
	for remainder != 0 && indexMap[remainder] == 0 {
		indexMap[remainder] = len(s)
		remainder *= 10
		s = append(s, '0'+byte(remainder/denominator))
		remainder %= denominator
	}
	if remainder > 0 { // 有循环节
		insertIndex := indexMap[remainder]
		s = append(s[:insertIndex], append([]byte{'('}, s[insertIndex:]...)...)
		s = append(s, ')')
	}
	return string(s)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

//------------------------Leetcode Problem 166------------------------
/*
 * https://leetcode.cn/problems/fraction-to-recurring-decimal/
 * 执行用时：0ms 在所有Go提交中击败了100.00%的用户
 * 占用内存：2MB 在所有Go提交中击败了74.36%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		numerator, _ := strconv.Atoi(scanner.Text())
		scanner.Scan()
		denominator, _ := strconv.Atoi(scanner.Text())
		Printf("Output: %v\n", fractionToDecimal(numerator, denominator))
	}
}
