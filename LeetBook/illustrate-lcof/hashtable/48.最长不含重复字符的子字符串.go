package main

import (
	"bufio"
	"os"
)

//请从字符串中找出一个最长的不包含重复字符的子字符串，计算该最长子字符串的长度。
func lengthOfLongestSubstring(s string) int {
	ans, idx := 0, -1
	mp := map[rune]int{}
	for i, x := range s {
		if j, y := mp[x]; y { //判断是否第二次出现
			idx = max(idx, j) //idx是该字符上次出现的位置
		}
		mp[x] = i             //更新字符的位置
		ans = max(ans, i-idx) //i-idx是该字符与上次的位置差
	}
	return ans
}

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		println(lengthOfLongestSubstring(input.Text()))
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

/*
 * 执行用时：4ms 在所有Go提交中击败了94.01%的用户
 * 占用内存：3.2MB 在所有Go提交中击败了22.92%的用户
**/
