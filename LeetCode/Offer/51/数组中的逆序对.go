package main

import (
	"bufio"
	. "fmt"
	"github.com/Don2025/GoCode/utils"
	"os"
	"sort"
)

// https://leetcode.cn/problems/shu-zu-zhong-de-ni-xu-dui-lcof/
// ------------------------剑指 Offer I Problem 51------------------------

type BIT struct {
	N    int
	Tree []int
}

func (b BIT) lowbit(x int) int { return x & (-x) }

func (b BIT) query(x int) int {
	var val int
	for x > 0 {
		val += b.Tree[x]
		x -= b.lowbit(x)
	}
	return val
}

func (b BIT) update(x int) {
	for x <= b.N {
		b.Tree[x]++
		x += b.lowbit(x)
	}
}

func reversePairs(nums []int) int {
	n := len(nums)
	tmp := make([]int, n)
	copy(tmp, nums)
	sort.Ints(tmp)
	for i := 0; i < n; i++ {
		nums[i] = sort.SearchInts(tmp, nums[i]) + 1
	}
	bit := BIT{n, make([]int, n+1)}
	var ans int
	for i := n - 1; i >= 0; i-- {
		ans += bit.query(nums[i] - 1)
		bit.update(nums[i])
	}
	return ans
}

// ------------------------剑指 Offer I Problem 51------------------------
/*
 * https://leetcode.cn/problems/shu-zu-zhong-de-ni-xu-dui-lcof/
 * 执行用时：136ms 在所有Go提交中击败了14.56%的用户
 * 占用内存：7.5MB 在所有Go提交中击败了78.44%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		nums := utils.StringToInts(scanner.Text())
		Printf("Output: %v\n", reversePairs(nums))
	}
}
