package main

import (
	"bufio"
	"container/list"
	"os"
	"strconv"
	"strings"
)

func validateStackSequences(pushed []int, popped []int) bool {
	stack := list.New()
	var cnt int
	for _, x := range pushed {
		stack.PushBack(x)
		for stack.Len() != 0 && stack.Back().Value.(int) == popped[cnt] {
			stack.Remove(stack.Back())
			cnt++
		}
	}
	return stack.Len() == 0
}

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		pushed := stringArrayToIntArray(strings.Fields(input.Text()))
		input.Scan()
		popped := stringArrayToIntArray(strings.Fields(input.Text()))
		println(validateStackSequences(pushed, popped))
	}
}

func stringArrayToIntArray(strings []string) []int {
	var arr []int
	for _, x := range strings {
		n, _ := strconv.Atoi(x)
		arr = append(arr, n)
	}
	return arr
}

/*
 * 执行用时：248ms 在所有Go提交中击败了55.83%的用户
 * 占用内存：8.7MB 在所有Go提交中击败了11.40%的用户
**/
