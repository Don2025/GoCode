package main

import (
	"bufio"
	. "fmt"
	"math"
	"os"
	"strings"
)

// https://leetcode.cn/problems/om3reC/
// ------------------------剑指 Offer II Problem 108------------------------
func ladderLength(beginWord string, endWord string, wordList []string) int {
	wordId := make(map[string]int)
	var graph [][]int
	addWord := func(word string) int {
		id, has := wordId[word]
		if !has {
			id = len(wordId)
			wordId[word] = id
			graph = append(graph, []int{})
		}
		return id
	}
	addEdge := func(word string) int {
		id1 := addWord(word)
		s := []byte(word)
		for i, b := range s {
			s[i] = '*'
			id2 := addWord(string(s))
			graph[id1] = append(graph[id1], id2)
			graph[id2] = append(graph[id2], id1)
			s[i] = b
		}
		return id1
	}
	for _, word := range wordList {
		addEdge(word)
	}
	beginId := addEdge(beginWord)
	endId, has := wordId[endWord]
	if !has {
		return 0
	}
	const inf int = math.MaxInt64
	dist := make([]int, len(wordId))
	for i := range dist {
		dist[i] = inf
	}
	dist[beginId] = 0
	queue := []int{beginId}
	for len(queue) > 0 {
		v := queue[0]
		queue = queue[1:]
		if v == endId {
			return dist[endId]>>1 + 1
		}
		for _, w := range graph[v] {
			if dist[w] == inf {
				dist[w] = dist[v] + 1
				queue = append(queue, w)
			}
		}
	}
	return 0
}

// ------------------------剑指 Offer II Problem 108------------------------
/*
 * https://leetcode.cn/problems/om3reC/
 * 执行用时：32ms 在所有Go提交中击败了89.02%的用户
 * 占用内存：9.6MB 在所有Go提交中击败了32.93%的用户
**/
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	Printf("Input a line of string(beginWord):")
	for scanner.Scan() {
		begin := scanner.Text()
		Printf("Input a line of string(endWord):")
		scanner.Scan()
		end := scanner.Text()
		Printf("Input a line of strings separated by space:")
		scanner.Scan()
		words := strings.Fields(scanner.Text())
		Printf("Output: %v\n", ladderLength(begin, end, words))
		Printf("Input a line of string(beginWord):")
	}
}
