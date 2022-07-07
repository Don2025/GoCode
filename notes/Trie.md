## [648. Replace Words](https://leetcode.cn/problems/replace-words/)

In English, we have a concept called **root**, which can be followed by some other word to form another longer word - let's call this word **successor**. For example, when the **root** `"an"` is followed by the **successor** word `"other"`, we can form a new word `"another"`.

Given a `dictionary` consisting of many **roots** and a `sentence` consisting of words separated by spaces, replace all the **successors** in the sentence with the **root** forming it. If a **successor** can be replaced by more than one **root**, replace it with the **root** that has **the shortest length**.

Return the `sentence` after the replacement.

##### Example 1:

```c
Input: dictionary = ["cat","bat","rat"], sentence = "the cattle was rattled by the battery"
Output: "the cat was rat by the bat"
```

##### Example 2:

```c
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

