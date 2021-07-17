package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func groupAnagrams(strs []string) [][]string {
	const maxn = 26
	mp := make(map[[maxn]int][]string)
	for _, s := range strs {
		cnt := [maxn]int{}
		for _, x := range s {
			cnt[x-'a']++
		}
		mp[cnt] = append(mp[cnt], s)
	}
	ans := make([][]string, 0, len(mp))
	for _, x := range mp {
		ans = append(ans, x)
	}
	return ans
}

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		ans := groupAnagrams(strings.Fields(input.Text()))
		for _, s := range ans {
			for _, x := range s {
				fmt.Printf("%s ", x)
			}
			fmt.Println()
		}
	}
}

/*
 * 执行用时：16ms 在所有Go提交中击败了100.00%的用户
 * 占用内存：8.5MB 在所有Go提交中击败了12.82%的用户
**/
