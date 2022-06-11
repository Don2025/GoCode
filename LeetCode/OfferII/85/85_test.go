package main

import (
	. "fmt"
	"testing"
)

type ques85 struct {
	argv85
	ans85 []string
}

type argv85 struct {
	one int
}

func Test85(t *testing.T) {
	examples := []ques85{
		{
			argv85{3},
			[]string{"((()))", "(()())", "(())()", "()(())", "()()()"},
		},
		{
			argv85{1},
			[]string{"()"},
		},
	}
	Println("------------------------剑指 Offer II Problem 85------------------------")
	for i, exam := range examples {
		input, expected := exam.argv85, exam.ans85
		Printf("Example %d:\n", i+1)
		Printf("Input: nums = %v\n", input.one)
		output := generateParenthesis(input.one)
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
