package main

//https://leetcode.cn/problems/US1pGT/
// ------------------------剑指 Offer II Problem 64------------------------

type MagicDictionary struct {
	LenMap map[int][]string
}

/** Initialize your data structure here. */
func Constructor() MagicDictionary {
	return MagicDictionary{map[int][]string{}}
}

func (md *MagicDictionary) BuildDict(dictionary []string) {
	for _, word := range dictionary {
		if _, ok := md.LenMap[len(word)]; ok {
			md.LenMap[len(word)] = append(md.LenMap[len(word)], word)
		} else {
			md.LenMap[len(word)] = []string{word}
		}
	}
}

func (md *MagicDictionary) Search(searchWord string) bool {
	if _, ok := md.LenMap[len(searchWord)]; !ok {
		return false
	}
	wordSlice := md.LenMap[len(searchWord)]
	for _, word := range wordSlice {
		var diffNum int
		for i := 0; i < len(word); i++ {
			if word[i] != searchWord[i] {
				diffNum++
			}
			if diffNum > 1 {
				break
			}
		}
		if diffNum == 1 {
			return true
		}
	}
	return false
}

/**
 * Your MagicDictionary object will be instantiated and called as such:
 * obj := Constructor();
 * obj.BuildDict(dictionary);
 * param_2 := obj.Search(searchWord);
 */

// ------------------------剑指 Offer II Problem 64------------------------
/*
 * https://leetcode.cn/problems/US1pGT/
 * 执行用时： 12ms 在所有Go提交中击败了 98.18%的用户
 * 占用内存： 6.8MB 在所有Go提交中击败了 95.45%的用户
**/
