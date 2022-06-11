package main

import (
	. "fmt"
	"github.com/Don2025/GoCode/utils"
	"testing"
)

type ques19 struct {
	argv19
	ans19 []int
}

type argv19 struct {
	one []int
	two int
}

func Test19(t *testing.T) {
	examples := []ques19{
		{
			argv19{[]int{1, 2, 3, 4, 5}, 2},
			[]int{1, 2, 3, 5},
		},
		{
			argv19{[]int{1}, 1},
			[]int{},
		},
		{
			argv19{[]int{1, 2}, 1},
			[]int{1},
		},
	}
	Println("------------------------Leetcode Problem 19------------------------")
	for i, exam := range examples {
		input, expected := exam.argv19, exam.ans19
		Printf("Example %d:\n", i+1)
		Printf("Input: head = %v, n = %d\n", input.one, input.two)
		output := utils.LinkListToInts(removeNthFromEnd(utils.IntsToLinkList(input.one), input.two))
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
