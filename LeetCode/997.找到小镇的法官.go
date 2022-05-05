package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func findJudge(n int, trust [][]int) int {
	type people struct{ Out, In int }
	p := make([]people, n)
	for _, x := range trust {
		p[x[0]-1].Out++ //x[0]是出度
		p[x[1]-1].In++  //x[1]是入度
	}
	for i := 0; i < n; i++ {
		if p[i].Out == 0 && p[i].In == n-1 {
			return i + 1
		}
	}
	return -1
}

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		n, _ := strconv.Atoi(input.Text())
		trust := make([][]int, n)
		for i := 0; i < n; i++ {
			input.Scan()
			trust[i] = stringArrayToIntArray(strings.Fields(input.Text()))
		}
		println(findJudge(n, trust))
	}
}

/*
 * 执行用时：96ms 在所有Go提交中击败了93.84%的用户
 * 占用内存：7.5MB 在所有Go提交中击败了85.87%的用户
**/
