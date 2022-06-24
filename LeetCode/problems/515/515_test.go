package main

import (
	. "fmt"
	"github.com/Don2025/GoCode/utils"
	"testing"
)

type ques515 struct {
	argv515
	ans515 []int
}

type argv515 struct {
	one string
}

func Test515(t *testing.T) {
	examples := []ques515{
		{
			argv515{"1 3 2 5 3 null 9"},
			[]int{1, 3, 9},
		},
		{
			argv515{"1 2 3"},
			[]int{1, 3},
		},
		{
			argv515{"0 -1"},
			[]int{0, -1},
		},
	}
	Println("------------------------Leetcode Problem 515------------------------")
	for i, exam := range examples {
		input, expected := exam.argv515, exam.ans515
		Printf("Example %d:\n", i+1)
		Printf("Input: root = %v\n", input.one)
		output := largestValues(utils.StringToTreeNode(input.one))
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
