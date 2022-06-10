package main

import (
	"bufio"
	"errors"
	"fmt"
	"github.com/Don2025/GoCode/utils"
	"os"
	"sort"
	"strconv"
	"testing"
)

func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums)
	n := len(nums)
	var quadruplets [][]int
	for i := 0; i < n-3 && nums[i]+nums[i+1]+nums[i+2]+nums[i+3] <= target; i++ {
		if i > 0 && nums[i] == nums[i-1] || nums[i]+nums[n-3]+nums[n-2]+nums[n-1] < target {
			continue
		}
		for j := i + 1; j < n-2 && nums[i]+nums[j]+nums[j+1]+nums[j+2] <= target; j++ {
			if j > i+1 && nums[j] == nums[j-1] || nums[i]+nums[j]+nums[n-2]+nums[n-1] < target {
				continue
			}
			for left, right := j+1, n-1; left < right; {
				if sum := nums[i] + nums[j] + nums[left] + nums[right]; sum == target {
					quadruplets = append(quadruplets, []int{nums[i], nums[j], nums[left], nums[right]})
					for left++; left < right && nums[left] == nums[left-1]; left++ {
					}
					for right--; left < right && nums[right] == nums[right+1]; right-- {
					}
				} else if sum < target {
					left++
				} else {
					right--
				}
			}
		}
	}
	return quadruplets
}

/*
 * https://leetcode.cn/problems/4sum/
 * 执行用时：4ms 在所有Go提交中击败了95.56%的用户
 * 占用内存：2.6MB 在所有Go提交中击败了77.67%的用户
**/

func main() {
	var t *testing.T = &testing.T{}
	Test18(t) // Use Example Testcases
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		nums := utils.StringToInts(scanner.Text())
		scanner.Scan()
		target, _ := strconv.Atoi(scanner.Text())
		fmt.Printf("%v\n", fourSum(nums, target))
	}
}

type ques18 struct {
	argv18
	ans18 [][]int
}

type argv18 struct {
	one []int
	two int
}

func Test18(t *testing.T) {
	examples := []ques18{
		{
			argv18{[]int{1, 0, -1, 0, -2, 2}, 0},
			[][]int{{-2, -1, 1, 2}, {-2, 0, 0, 2}, {-1, 0, 0, 1}},
		},
		{
			argv18{[]int{2, 2, 2, 2, 2}, 8},
			[][]int{{2, 2, 2, 2}},
		},
	}
	fmt.Println("------------------------Leetcode Problem 18------------------------")
	for i, exam := range examples {
		input, expected := exam.argv18, exam.ans18
		fmt.Printf("Example %d:\n", i+1)
		fmt.Printf("Input: nums = %v, target = %d\n", input.one, input.two)
		output := fourSum(input.one, input.two)
		fmt.Printf("Output: %v\n", output)
		fmt.Printf("Expected: %v\n", expected)
		if fmt.Sprintf("%v", output) == fmt.Sprintf("%v", expected) {
			fmt.Printf("There's not much of a difference.\n")
			fmt.Println("-------------------------------------------------------------------")
		} else {
			err := fmt.Sprintf("Example %d went wrong!\n", i+1)
			panic(errors.New(err))
		}
	}
	fmt.Println("You can input some customize test case right now.")
}
