package main

import (
	"bufio"
	. "fmt"
	"os"
	"strconv"
)

// https://leetcode.cn/problems/maximize-the-confusion-of-an-exam/
//------------------------Leetcode Problem 2024------------------------
func maxConsecutiveAnswers(answerKey string, k int) int {
	m := make(map[byte]int)
	var ans int
	for i, j := 0, 0; j < len(answerKey); j++ {
		m[answerKey[j]]++
		if min(m['T'], m['F']) > k {
			m[answerKey[i]]--
			i++
		}
		ans = max(ans, j-i+1)
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

//------------------------Leetcode Problem 2024------------------------
/*
 * https://leetcode.cn/problems/maximize-the-confusion-of-an-exam/
 * 执行用时：44ms 在所有Go提交中击败了6.45%的用户
 * 占用内存：5.1MB 在所有Go提交中击败了29.03%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		s := scanner.Text()
		scanner.Scan()
		k, _ := strconv.Atoi(scanner.Text())
		Printf("Output: %v\n", maxConsecutiveAnswers(s, k))
	}
}
