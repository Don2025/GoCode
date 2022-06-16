package main

import (
	. "fmt"
	"testing"
)

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
	Println("------------------------Leetcode Problem 18------------------------")
	for i, exam := range examples {
		input, expected := exam.argv18, exam.ans18
		Printf("Example %d:\n", i+1)
		Printf("Input: nums = %v, target = %d\n", input.one, input.two)
		output := fourSum(input.one, input.two)
		Printf("Output: %v\n", output)
		Printf("Expected: %v\n", expected)
		if Sprintf("%v", output) == Sprintf("%v", expected) {
			Printf("There's not much of a difference.\n")
			Println("-------------------------------------------------------------------")
		} else {
			err := Sprintf("Example %d went wrong!\n", i+1)
			t.Errorf(err) // Fail
			// t.Fatalf(err) // FailNow
			// panic(errors.New(err))
		}
	}
	Println("You can input some customize test case right now.")
}
