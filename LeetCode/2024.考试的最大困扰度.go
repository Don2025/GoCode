package main

import (
	"bufio"
	"os"
	"strconv"
)

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

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		s := input.Text()
		input.Scan()
		k, _ := strconv.Atoi(input.Text())
		println(maxConsecutiveAnswers(s, k))
	}
}

/*
 * 执行用时：44ms 在所有Go提交中击败了6.45%的用户
 * 占用内存：5.1MB 在所有Go提交中击败了29.03%的用户
**/
