package main

import (
	"bufio"
	. "fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

// https://leetcode.cn/problems/top-k-frequent-words/
//------------------------Leetcode Problem 692------------------------
func topKFrequent(words []string, k int) []string {
	cnt := map[string]int{}
	for _, x := range words {
		cnt[x]++
	}
	uniqueWords := make([]string, 0, len(cnt))
	for x := range cnt {
		uniqueWords = append(uniqueWords, x)
	}
	sort.Slice(uniqueWords, func(i, j int) bool {
		s, t := uniqueWords[i], uniqueWords[j]
		return cnt[s] > cnt[t] || cnt[s] == cnt[t] && s < t
	})
	return uniqueWords[:k]
}

//------------------------Leetcode Problem 692------------------------
/*
 * https://leetcode.cn/problems/top-k-frequent-words/
 * 执行用时：8ms 在所有Go提交中击败了81.34%的用户
 * 占用内存：4.2MB 在所有Go提交中击败了98.51%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		strArr := strings.Fields(scanner.Text())
		scanner.Scan()
		k, _ := strconv.Atoi(scanner.Text())
		Printf("Output: %v\n", topKFrequent(strArr, k))
	}
}
