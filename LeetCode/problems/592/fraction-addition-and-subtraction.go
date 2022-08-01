package main

import (
	"bufio"
	. "fmt"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
)

// https://leetcode.cn/problems/fraction-addition-and-subtraction/
//------------------------Leetcode Problem 592------------------------
func fractionAddition(expression string) string {
	regex, _ := regexp.Compile("[+-]?\\d+")
	expressionArr := regex.FindAllString(expression, -1)
	nums := make([]int, len(expressionArr))
	for i := 0; i < len(expressionArr); i++ {
		if (i & 1) == 1 {
			denominator, _ := strconv.Atoi(expressionArr[i])
			nums[i] = denominator
		} else {
			numerator, _ := strconv.Atoi(expressionArr[i])
			nums[i] = numerator
		}
	}
	numerator, denominator := 0, 1
	for i := 0; i < len(expressionArr); i++ {
		if (i & 1) == 1 {
			denominator *= nums[i]
		}
	}
	for i := 0; i < len(expressionArr); i++ {
		if (i & 1) == 0 {
			numerator += nums[i] * denominator / nums[i+1]
		}
	}
	gcd := func(a, b int) int {
		if b == 0 {
			return a
		}
		for a%b != 0 {
			a, b = b, a%b
		}
		return b
	}
	common := int(math.Abs(float64(gcd(denominator, numerator))))
	return strings.Join([]string{strconv.Itoa(numerator / common), "/", strconv.Itoa(denominator / common)}, "")
}

//------------------------Leetcode Problem 592------------------------
/*
 * https://leetcode.cn/problems/fraction-addition-and-subtraction/
 * 执行用时：0ms 在所有Go提交中击败了100.00%的用户
 * 占用内存：2.4MB 在所有Go提交中击败了5.37%的用户
**/
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		Printf("Output: %v\n", fractionAddition(scanner.Text()))
	}
}
