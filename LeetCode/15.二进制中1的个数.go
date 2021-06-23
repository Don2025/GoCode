package main

import (
	"bufio"
	"math/bits"
	"os"
	"strconv"
)

func hammingWeight(num uint32) int {
	return bits.OnesCount32(num)
}

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		n, _ := strconv.Atoi(input.Text())
		println(hammingWeight(uint32(n)))
	}
}

/*
 * 执行用时：0ms 在所有Go提交中击败了100.00%的用户
 * 占用内存：1.9MB 在所有Go提交中击败了61.94%的用户
**/
