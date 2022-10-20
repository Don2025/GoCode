package main

import (
	"bufio"
	. "fmt"
	"os"
)

// https://leetcode.cn/problems/k-th-symbol-in-grammar/
//------------------------Leetcode Problem 779------------------------
func kthGrammar(n, k int) int {
	if n == 1 {
		return 0
	}
	return k&1 ^ 1 ^ kthGrammar(n-1, (k+1)/2)
}

//------------------------Leetcode Problem 779------------------------
/*
 * https://leetcode.cn/problems/k-th-symbol-in-grammar/
 * 执行用时：0ms 在所有Go提交中击败了100.00%的用户
 * 占用内存：1.9MB 在所有Go提交中击败了34.09%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		var n, k int
		Sscanf(scanner.Text(), "%d %d", &n, &k)
		Printf("Output: %d\n", kthGrammar(n, k))
	}
}
