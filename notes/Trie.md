## [648. Replace Words](https://leetcode.cn/problems/replace-words/)

In English, we have a concept called **root**, which can be followed by some other word to form another longer word - let's call this word **successor**. For example, when the **root** `"an"` is followed by the **successor** word `"other"`, we can form a new word `"another"`.

Given a `dictionary` consisting of many **roots** and a `sentence` consisting of words separated by spaces, replace all the **successors** in the sentence with the **root** forming it. If a **successor** can be replaced by more than one **root**, replace it with the **root** that has **the shortest length**.

Return the `sentence` after the replacement.

##### Example 1:

```go
Input: dictionary = ["cat","bat","rat"], sentence = "the cattle was rattled by the battery"
Output: "the cat was rat by the bat"
```

##### Example 2:

```go
Input: dictionary = ["a","b","c"], sentence = "aadsfasf absbs bbab cadsfafs"
Output: "a a b c"
```

##### Solution:

```go
func replaceWords(dictionary []string, sentence string) string {
	root := &Trie{}
	for _, d := range dictionary {
		root.Insert(d)
	}
	words := strings.Fields(sentence)
	for i, word := range words {
		if w, ok := root.Search(word); ok {
			words[i] = w
		}
	}
	return strings.Join(words, " ")
}

type Trie struct {
	Children [26]*Trie
	IsWord   bool
}

func (root *Trie) Insert(word string) {
	node := root
	for _, c := range word {
		idx := c - 'a'
		if node.Children[idx] == nil {
			node.Children[idx] = &Trie{}
		}
		node = node.Children[idx]
	}
	node.IsWord = true
}

func (root *Trie) Search(word string) (string, bool) {
	node := root
	for i, c := range word {
		if node.IsWord {
			if _, ok := root.Search(word[i:]); ok {
				return word[:i], true
			}
		}
		idx := c - 'a'
		if node.Children[idx] == nil {
			break
		}
		node = node.Children[idx]
	}
	return word, node.IsWord
}
```

------

## [720. Longest Word in Dictionary](https://leetcode.cn/problems/longest-word-in-dictionary/)

Given an array of strings words representing an English Dictionary, return the longest word in words that can be built one character at a time by other words in words.

If there is more than one possible answer, return the longest word with the smallest lexicographical order. If there is no answer, return the empty string.

Note that the word should be built from left to right with each additional character being added to the end of a previous word. 

##### Example 1:

```go
Input: words = ["w","wo","wor","worl","world"]
Output: "world"
Explanation: The word "world" can be built one character at a time by "w", "wo", "wor", and "worl".
```

##### Example 2:

```go
Input: words = ["a","banana","app","appl","ap","apply","apple"]
Output: "apple"
Explanation: Both "apply" and "apple" can be built from other words in the dictionary. However, "apple" is lexicographically smaller than "apply".
```

**Constraints:**

- `1 <= words.length <= 1000`
- `1 <= words[i].length <= 30`
- `words[i]` consists of lowercase English letters.

```go
type Trie struct {
	Children [26]*Trie
	IsWord    bool
}

func (t *Trie) Insert(word string) {
	node := t
	for _, x := range word {
		x -= 'a'
		if node.Children[x] == nil {
			node.Children[x] = &Trie{}
		}
		node = node.Children[x]
	}
	node.IsWord = true
}

func (t *Trie) Search(word string) bool {
	node := t
	for _, x := range word {
		x -= 'a'
		if node.Children[x] == nil || !node.Children[x].IsWord {
			return false
		}
		node = node.Children[x]
	}
	return true
}

func longestWord(words []string) string {
	t := &Trie{}
	for _, word := range words {
		t.Insert(word)
	}
	var ans string
	for _, word := range words {
		if t.Search(word) && (len(word) > len(ans) || len(word) == len(ans) && word < ans) {
			ans = word
		}
	}
	return ans
}
```

------



```go
// https://leetcode.cn/problems/maximum-xor-with-an-element-from-array/
//------------------------Leetcode Problem 1707------------------------

const L = 30

type trieNode struct {
	children [2]*trieNode
}

func maximizeXor(nums []int, queries [][]int) (ans []int) {
	sort.Ints(nums)
	for i := range queries {
		queries[i] = append(queries[i], i)
	}
	sort.Slice(queries, func(i, j int) bool {
		return queries[i][1] < queries[j][1]
	})
	ans = make([]int, len(queries))
	t := &trieNode{}
	idx, n := 0, len(nums)
	for _, q := range queries {
		x, m, qid := q[0], q[1], q[2]
		for idx < n && nums[idx] <= m {
			t.insert(nums[idx])
			idx++
		}
		if idx == 0 {
			ans[qid] = -1
		} else {
			ans[qid] = t.getMaxXor(x)
		}
	}
	return
}

func (t *trieNode) insert(val int) {
	node := t
	for i := L - 1; i >= 0; i-- {
		bit := val >> i & 1
		if node.children[bit] == nil {
			node.children[bit] = &trieNode{}
		}
		node = node.children[bit]
	}
}

func (t *trieNode) getMaxXor(val int) (ans int) {
	node := t
	for i := L - 1; i >= 0; i-- {
		bit := val >> i & 1
		if node.children[bit^1] != nil {
			ans |= 1 << i
			bit ^= 1
		}
		node = node.children[bit]
	}
	return
}
```

