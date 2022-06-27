package main

import (
	"bufio"
	. "fmt"
	"os"
	"strconv"
)

// https://leetcode.cn/problems/reverse-bits/
//------------------------Leetcode Problem 190------------------------
func reverseBits(num uint32) uint32 {
	var ans uint32
	for i := 32; i > 0; i-- {
		ans <<= 1
		ans += num & 1
		num >>= 1
	}
	return ans
}

//------------------------Leetcode Problem 190------------------------
/*
 * https://leetcode.cn/problems/reverse-bits/
 * 执行用时：0ms 在所有Go提交中击败了100.00%的用户
 * 占用内存：2.4MB 在所有Go提交中击败了76.98%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		num, _ := strconv.ParseUint(scanner.Text(), 2, 32)
		ansInt := reverseBits(uint32(num))
		ansUint32 := strconv.FormatUint(uint64(ansInt), 2)
		ansUint32 = Sprintf("%0*v", 32, ansUint32)
		Printf("Output: %v (%v)\n", ansInt, ansUint32)
	}
}
