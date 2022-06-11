package main

import (
	"errors"
	. "fmt"
	"testing"
)

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
