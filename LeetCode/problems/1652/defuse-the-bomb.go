package main

import (
	"bufio"
	. "fmt"
	"github.com/Don2025/GoCode/utils"
	"os"
	"strconv"
)

// https://leetcode.cn/problems/defuse-the-bomb/
//------------------------Leetcode Problem 1652------------------------
func decrypt(code []int, k int) []int {
	n := len(code)
	ans := make([]int, n)
	if k == 0 {
		return ans
	}
	code = append(code, code...)
	l, r := 1, k
	if k < 0 {
		l, r = n+k, n-1
	}
	sum := 0
	for _, v := range code[l : r+1] {
		sum += v
	}
	for i := range ans {
		ans[i] = sum
		sum -= code[l]
		sum += code[r+1]
		l, r = l+1, r+1
	}
	return ans
}

// ------------------------Leetcode Problem1652------------------------
/*
 * https://leetcode.cn/problems/defuse-the-bomb/
 * 执行用时：0ms 在所有Go提交中击败了100.00%的用户
 * 占用内存：2.3MB 在所有Go提交中击败了21.16%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		code := utils.StringToInts(scanner.Text())
		scanner.Scan()
		k, _ := strconv.Atoi(scanner.Text())
		Printf("Output: %d\n", decrypt(code, k))
	}
}
