package main

// https://leetcode.cn/problems/prefix-and-suffix-search/
//------------------------Leetcode Problem 745------------------------

type Pair struct {
	Prefix string
	Suffix string
}

type WordFilter map[Pair]int

func Constructor(words []string) WordFilter {
	wf := WordFilter{}
	for idx, word := range words {
		for i := 1; i <= len(word); i++ {
			for j := 0; j < len(word); j++ {
				wf[Pair{word[:i], word[j:]}] = idx
			}
		}
	}
	return wf
}

func (wf WordFilter) F(pref string, suff string) int {
	if i, ok := wf[Pair{pref, suff}]; ok {
		return i
	}
	return -1
}

/**
 * Your WordFilter object will be instantiated and called as such:
 * obj := Constructor(words);
 * param_1 := obj.F(pref,suff);
 */
//------------------------Leetcode Problem 745------------------------
/*
 * https://leetcode.cn/problems/prefix-and-suffix-search/
 * 执行用时：664ms 在所有Go提交中击败了20.00%的用户
 * 占用内存：42.7MB 在所有Go提交中击败了52.00%的用户
**/
