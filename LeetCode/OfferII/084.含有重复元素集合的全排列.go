package main

import (
	"bufio"
	"errors"
	. "fmt"
	"github.com/Don2025/GoCode/utils"
	"os"
	"sort"
	"testing"
)

// https://leetcode.cn/problems/7p8L0Z/
// ------------------------剑指 Offer II Problem 84------------------------
func permuteUnique(nums []int) (ans [][]int) {
	sort.Ints(nums)
	n := len(nums)
	vis := make([]bool, n)
	var arr []int
	var backtrack func(int)
	backtrack = func(idx int) {
		if idx == n {
			ans = append(ans, append([]int(nil), arr...))
			return
		}
		for i, x := range nums {
			if vis[i] || i > 0 && !vis[i-1] && x == nums[i-1] {
				continue
			}
			arr = append(arr, x)
			vis[i] = true
			backtrack(idx + 1)
			vis[i] = false
			arr = arr[:len(arr)-1]
		}
	}
	backtrack(0)
	return
}

//------------------------剑指 Offer II Problem 84------------------------
/*
 * https://leetcode.cn/problems/7p8L0Z/
 * 执行用时：0ms 在所有Go提交中击败了100.00%的用户
 * 占用内存：3.6MB 在所有Go提交中击败了98.31%的用户
**/
func main() {
	var t *testing.T = &testing.T{}
	Test84(t) // Use Example Testcases
	scanner := bufio.NewScanner(os.Stdin)
	Printf("Input a line of numbers separated by space:")
	for scanner.Scan() {
		Printf("Output: %v\n", permuteUnique(utils.StringToInts(scanner.Text())))
		Printf("Input a line of numbers separated by space:")
	}
}

type ques84 struct {
	argv84
	ans84 [][]int
}

type argv84 struct {
	one []int
}

func Test84(t *testing.T) {
	examples := []ques84{
		{
			argv84{[]int{1, 1, 2}},
			[][]int{{1, 1, 2}, {1, 2, 1}, {2, 1, 1}},
		},
		{
			argv84{[]int{1, 2, 3}},
			[][]int{{1, 2, 3}, {1, 3, 2}, {2, 1, 3}, {2, 3, 1}, {3, 1, 2}, {3, 2, 1}},
		},
	}
	Println("------------------------剑指 Offer II Problem 84------------------------")
	for i, exam := range examples {
		input, expected := exam.argv84, exam.ans84
		Printf("Example %d:\n", i+1)
		Printf("Input: nums = %v\n", input.one)
		output := permuteUnique(input.one)
		Printf("Output: %v\n", output)
		Printf("Expected: %v\n", expected)
		if Sprintf("%v", output) == Sprintf("%v", expected) {
			Printf("There's not much of a difference.\n")
			Println("-------------------------------------------------------------------")
		} else {
			err := Sprintf("Example %d went wrong!\n", i+1)
			panic(errors.New(err))
		}
	}
	Println("You can input some customize test case right now.")
}
