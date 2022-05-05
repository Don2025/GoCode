package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func groupAnagrams(strs []string) [][]string {
	m := make(map[string][]string)
	for _, str := range strs {
		s := []byte(str)
		sort.Slice(s, func(i, j int) bool { return s[i] < s[j] })
		sortedStr := string(s)
		m[sortedStr] = append(m[sortedStr], str)
	}
	ans := make([][]string, 0, len(m))
	for _, x := range m {
		ans = append(ans, x)
	}
	return ans
}

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		arr := strings.Fields(input.Text())
		ans := groupAnagrams(arr)
		for _, s := range ans {
			fmt.Printf("%v ", s)
		}
		fmt.Println()
	}
}

/*
 * 执行用时：24ms 在所有Go提交中击败了53.85%的用户
 * 占用内存：8.2MB 在所有Go提交中击败了60.982%的用户
**/
