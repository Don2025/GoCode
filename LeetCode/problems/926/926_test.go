package main

import (
	. "fmt"
	"testing"
)

type ques926 struct {
	argv926
	ans926 int
}

type argv926 struct {
	one string
}

func Test926(t *testing.T) {
	examples := []ques926{
		{
			argv926{"00110"},
			1,
		},
		{
			argv926{"010110"},
			2,
		},
		{
			argv926{"00011000"},
			2,
		},
	}
	Println("------------------------Leetcode Problem 926------------------------")
	for i, exam := range examples {
		input, expected := exam.argv926, exam.ans926
		Printf("Example %d:\n", i+1)
		Printf("Input: head = %s\n", input.one)
		output := minFlipsMonoIncr(input.one)
		Printf("Output: %v\n", output)
		Printf("Expected: %v\n", expected)
		if output == expected {
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
