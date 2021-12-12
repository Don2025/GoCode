package main

import (
	"bufio"
	"os"
)

func countPoints(rings string) int {
	a := [10][3]bool{}
	for i := 0; i < len(rings); i++ {
		if rings[i] == 'R' {
			a[rings[i+1]-'0'][0] = true
		} else if rings[i] == 'G' {
			a[rings[i+1]-'0'][1] = true
		} else if rings[i] == 'B' {
			a[rings[i+1]-'0'][2] = true
		}
	}
	var cnt int
	for i := range a {
		if a[i][0] && a[i][1] && a[i][2] {
			cnt++
		}
	}
	return cnt
}

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		println(countPoints(input.Text()))
	}
}
