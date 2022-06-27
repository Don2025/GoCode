package main

import (
	"bufio"
	. "fmt"
	"os"
)

// https://leetcode.cn/problems/repeated-dna-sequences/
//------------------------Leetcode Problem 187------------------------
func findRepeatedDnaSequences(s string) []string {
	const L = 10
	var ans []string
	cnt := map[string]int{}
	for i := 0; i <= len(s)-L; i++ {
		sub := s[i : i+L]
		cnt[sub]++
		if cnt[sub] == 2 {
			ans = append(ans, sub)
		}
	}
	return ans
}

//------------------------Leetcode Problem 187------------------------
/*
 * https://leetcode.cn/problems/repeated-dna-sequences/
 * 执行用时：8ms 在所有Go提交中击败了99.28%的用户
 * 占用内存：9.2MB 在所有Go提交中击败了16.71%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		Printf("%v\n", findRepeatedDnaSequences(scanner.Text()))
	}
}
