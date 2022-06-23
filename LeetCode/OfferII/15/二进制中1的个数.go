package main

import (
	"bufio"
	. "fmt"
	"math/bits"
	"os"
	"strconv"
)

// ------------------------剑指 Offer II Problem 15------------------------
// https://leetcode.cn/problems/er-jin-zhi-zhong-1de-ge-shu-lcof/
func hammingWeight(num uint32) int {
	return bits.OnesCount32(num)
}

// ------------------------剑指 Offer II Problem 15------------------------
/*
 * https://leetcode.cn/problems/er-jin-zhi-zhong-1de-ge-shu-lcof/
 * 执行用时：0ms 在所有Go提交中击败了100.00%的用户
 * 占用内存：1.9MB 在所有Go提交中击败了61.94%的用户
**/
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		n, _ := strconv.Atoi(scanner.Text())
		Printf("%v\n", hammingWeight(uint32(n)))
	}
}
