package main

import (
	"bufio"
	"bytes"
	"os"
	"sort"
)

func frequencySort(s string) string {
	mp := make(map[rune]int)
	for _, x := range s {
		mp[x]++
	}
	type pair struct {
		First  rune
		Second int
	}
	pairs := make([]pair, 0, len(mp))
	for i, x := range mp {
		pairs = append(pairs, pair{i, x})
	}
	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].Second > pairs[j].Second
	})
	ans := make([]byte, 0, len(s))
	for _, x := range pairs {
		ans = append(ans, bytes.Repeat([]byte{byte(x.First)}, x.Second)...)
	}
	return string(ans)
}

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		println(frequencySort(input.Text()))
	}
}

/*
 * 执行用时：4ms 在所有Go提交中击败了96.12%的用户
 * 占用内存：4.8MB 在所有Go提交中击败了88.37%的用户
**/
