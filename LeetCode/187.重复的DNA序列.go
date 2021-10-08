package main

import (
	"bufio"
	"os"
)

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

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		ans := findRepeatedDnaSequences(input.Text())
		for _, x := range ans {
			println(x)
		}
	}
}
