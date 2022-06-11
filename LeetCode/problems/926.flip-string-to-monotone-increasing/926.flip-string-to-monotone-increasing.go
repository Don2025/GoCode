package main

import (
	"bufio"
	"errors"
	. "fmt"
	"os"
	"testing"
)

// https://leetcode.cn/problems/flip-string-to-monotone-increasing/
//------------------------Leetcode Problem 926------------------------
func minFlipsMonoIncr(s string) int {
	var dp0, dp1 int
	for _, ch := range s {
		dp0New, dp1New := dp0, min(dp0, dp1)
		if ch == '1' {
			dp0New++
		} else {
			dp1New++
		}
		dp0, dp1 = dp0New, dp1New
	}
	return min(dp0, dp1)
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

//------------------------Leetcode Problem 926------------------------
func main() {
	var t *testing.T = &testing.T{}
	Test926(t) // Use Example Testcases
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		Println(minFlipsMonoIncr(scanner.Text()))
	}
}

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
			panic(errors.New(err))
		}
	}
	Println("You can input some customize test case right now.")
}
