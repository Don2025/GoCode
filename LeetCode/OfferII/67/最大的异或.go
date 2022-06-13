package main

import (
	"bufio"
	. "fmt"
	"github.com/Don2025/GoCode/utils"
	"os"
)

// https://leetcode.cn/problems/ms70jA/
// ------------------------剑指 Offer II Problem 67------------------------
const highBit = 30

type Trie struct {
	Left, Right *Trie
}

func (t *Trie) add(num int) {
	cur := t
	for i := highBit; i >= 0; i-- {
		bit := num >> i & 1
		if bit == 0 {
			if cur.Left == nil {
				cur.Left = &Trie{}
			}
			cur = cur.Left
		} else {
			if cur.Right == nil {
				cur.Right = &Trie{}
			}
			cur = cur.Right
		}
	}
}

func (t *Trie) check(num int) (x int) {
	cur := t
	for i := highBit; i >= 0; i-- {
		bit := num >> i & 1
		if bit == 0 {
			if cur.Right != nil {
				cur = cur.Right
				x = x*2 + 1
			} else {
				cur = cur.Left
				x = x * 2
			}
		} else {
			if cur.Left != nil {
				cur = cur.Left
				x = x*2 + 1
			} else {
				cur = cur.Right
				x = x * 2
			}
		}
	}
	return
}

func findMaximumXOR(nums []int) (x int) {
	root := &Trie{}
	for i := 1; i < len(nums); i++ {
		root.add(nums[i-1])
		x = max(x, root.check(nums[i]))
	}
	return
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// ------------------------剑指 Offer II Problem 67------------------------
/*
 * https://leetcode.cn/problems/ms70jA/
 * 执行用时：132ms 在所有Go提交中击败了70.59%的用户
 * 占用内存：9.2MB 在所有Go提交中击败了86.56%的用户
**/
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	Printf("Input a line of numbers separated by space:")
	for scanner.Scan() {
		nums := utils.StringToInts(scanner.Text())
		Printf("Output: %v\n", findMaximumXOR(nums))
		Printf("Input a line of numbers separated by space:")
	}
}
