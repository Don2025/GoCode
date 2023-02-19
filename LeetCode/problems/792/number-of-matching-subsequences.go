package main

import (
	"bufio"
	. "fmt"
	"os"
	"strings"
)

// https://leetcode.cn/problems/number-of-matching-subsequences/
//------------------------Leetcode Problem 792------------------------
func numMatchingSubseq(s string, words []string) (cnt int) {
	type pair struct{ i, j int }
	queues := make([][]pair, 26)
	for i, w := range words {
		queues[w[0]-'a'] = append(queues[w[0]-'a'], pair{i, 0})
	}
	for _, c := range s {
		q := queues[c-'a']
		queues[c-'a'] = nil
		for _, p := range q {
			p.j++
			if p.j == len(words[p.i]) {
				cnt++
			} else {
				queues[words[p.i][p.j]-'a'] = append(queues[words[p.i][p.j]-'a'], p)
			}
		}
	}
	return
}

//------------------------Leetcode Problem 792------------------------
/*
 * https://leetcode.cn/problems/number-of-matching-subsequences/
 * 执行用时：48ms 在所有Go提交中击败了82.35%的用户
 * 占用内存：9.2MB 在所有Go提交中击败了27.94%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		s := scanner.Text()
		words := strings.Fields(scanner.Text())
		Printf("Output: %v\n", numMatchingSubseq(s, words))
	}
}
