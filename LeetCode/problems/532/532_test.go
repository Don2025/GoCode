package main

import (
	. "fmt"
	"testing"
)

type ques532 struct {
	argv532
	ans532 int
}

type argv532 struct {
	one []int
	two int
}

func Test532(t *testing.T) {
	examples := []ques532{
		{
			argv532{[]int{3, 1, 4, 1, 5}, 2},
			2,
		},
		{
			argv532{[]int{1, 2, 3, 4, 5}, 1},
			4,
		},
		{
			argv532{[]int{1, 3, 1, 5, 4}, 0},
			1,
		},
	}
	Println("------------------------Leetcode Problem 532------------------------")
	for i, exam := range examples {
		input, expected := exam.argv532, exam.ans532
		Printf("Example %d:\n", i+1)
		Printf("Input: head = %v, n = %d\n", input.one, input.two)
		output := findPairs(input.one, input.two)
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
