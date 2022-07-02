package main

import (
	"bufio"
	. "fmt"
	"os"
	"strconv"
)

// https://leetcode.cn/problems/2-keys-keyboard/
//------------------------Leetcode Problem 650------------------------
func minSteps(n int) int {
	var ans int
	for i := 2; i*i <= n; i++ {
		for n%i == 0 {
			n /= i
			ans += i
		}
	}
	if n > 1 {
		ans += n
	}
	return ans
}

//------------------------Leetcode Problem 650------------------------
/*
 * https://leetcode.cn/problems/2-keys-keyboard/
 * 执行用时：0-ms 在所有Go提交中击败了100.00%的用户
 * 占用内存：1.9MB 在所有Go提交中击败了90.24%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		n, _ := strconv.Atoi(scanner.Text())
		Printf("Output: %v\n", minSteps(n))
	}
}
