package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func findAnagrams(s string, p string) []int {
	var ans []int
	if len(s) < len(p) {
		return ans
	}
	freq_s, freq_p := [26]int{}, [26]int{}
	for i, x := range p {
		freq_p[x-'a']++
		freq_s[s[i]-'a']++
	}
	if freq_s == freq_p {
		ans = append(ans, 0)
	}
	for i, x := range s[:len(s)-len(p)] {
		freq_s[x-'a']--
		freq_s[s[i+len(p)]-'a']++
		if freq_s == freq_p {
			ans = append(ans, i+1)
		}
	}
	return ans
}

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		arr := strings.Fields(input.Text())
		s, p := arr[0], arr[1]
		ans := findAnagrams(s, p)
		for _, x := range ans {
			fmt.Printf("%d ", x)
		}
		fmt.Println()
	}
}

/*
 * 执行用时：4ms 在所有Go提交中击败了94.79%的用户
 * 占用内存：5MB 在所有Go提交中击败了79.68%的用户
**/
