package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

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

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		strArr := strings.Fields(input.Text())
		input.Scan()
		k, _ := strconv.Atoi(input.Text())
		ansArr := topKFrequent(strArr, k)
		for _, x := range ansArr {
			fmt.Printf("%s ", x)
		}
		println()
	}
}

/*
 * 执行用时：8ms 在所有Go提交中击败了81.34%的用户
 * 占用内存：4.2MB 在所有Go提交中击败了98.51%的用户
**/
