package main

import (
	"bufio"
	. "fmt"
	"os"
	"strconv"
)

// https://leetcode.cn/problems/find-the-winner-of-the-circular-game/
//------------------------Leetcode Problem 1823------------------------
func findTheWinner(n int, k int) int {
	var p int
	for i := 2; i <= n; i++ {
		p = (p + k) % i
	}
	return p + 1
}

//------------------------Leetcode Problem 1823------------------------
/*
 * https://leetcode.cn/problems/find-the-winner-of-the-circular-game/
 * 执行用时：0ms 在所有Go提交中击败了100.00%的用户
 * 占用内存：1.9MB 在所有Go提交中击败了79.19%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		n, _ := strconv.Atoi(scanner.Text())
		scanner.Scan()
		k, _ := strconv.Atoi(scanner.Text())
		Printf("Output: %v\n", findTheWinner(n, k))
	}
}
