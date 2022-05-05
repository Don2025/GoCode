package main

import (
	"bufio"
	"math"
	"os"
	"strconv"
	"strings"
)

func verifyPostorder(postorder []int) bool {
	var stack []int
	root := math.MaxInt32
	for i := len(postorder) - 1; i >= 0; i-- {
		if postorder[i] > root {
			return false
		}
		for len(stack) > 0 && stack[len(stack)-1] > postorder[i] {
			root = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, postorder[i])
	}
	return true
}

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		postorder := stringArrayToIntArray(strings.Fields(input.Text()))
		println(verifyPostorder(postorder))
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
 * 执行用时：0ms 在所有Go提交中击败了100.00%的用户
 * 占用内存：1.9MB 在所有Go提交中击败了26.74%的用户
**/
