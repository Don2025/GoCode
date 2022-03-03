package main

import (
	"bufio"
	"math/bits"
	"os"
	"strconv"
)

/*
注意到 right ≤ 10的6次方 < 2的20次方
因此二进制中1的个数不会超过19，而不超过19的质数只有 2, 3, 5, 7, 11, 13, 17, 19
我们可以用一个二进制数mask=dec(665772)=bin(10100010100010101100)来存储这些质数
其中mask二进制的从低到高的第 i 位为 1 表示 i 是质数，为 0 表示 i 不是质数
设整数x的二进制中1的个数为c 若mask按位与 2^c不为0则说明 c 是一个质数
*/
func countPrimeSetBits(left, right int) int {
	var ans int
	for i := left; i <= right; i++ {
		if 1<<bits.OnesCount(uint(i))&665772 != 0 {
			ans++
		}
	}
	return ans
}

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		left, _ := strconv.Atoi(input.Text())
		input.Scan()
		right, _ := strconv.Atoi(input.Text())
		println(countPrimeSetBits(left, right))
	}
}

/*
 * 执行用时：0ms 在所有Go提交中击败了100.00%的用户
 * 占用内存：1.8MB 在所有Go提交中击败了93.94%的用户
**/
