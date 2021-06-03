package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

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

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		nums := stringArrayToIntArray(strings.Fields(input.Text()))
		input.Scan()
		n, _ := strconv.Atoi(input.Text())
		queries := make([][]int, n)
		for i := 0; i < n; i++ {
			input.Scan()
			queries[i] = stringArrayToIntArray(strings.Fields(input.Text()))
		}
		ans := maximizeXor(nums, queries)
		for _, x := range ans {
			fmt.Printf("%d ", x)
		}
		println()
	}
}

func stringArrayToIntArray(strings []string) []int {
	var arr []int
	for _, x := range strings {
		n, _ := strconv.Atoi(x)
		arr = append(arr, n)
	}
	return arr
}

/*
 * 执行用时：612ms 在所有Go提交中击败了16.67%的用户
 * 占用内存：44.6MB 在所有Go提交中击败了50.00%的用户
**/
