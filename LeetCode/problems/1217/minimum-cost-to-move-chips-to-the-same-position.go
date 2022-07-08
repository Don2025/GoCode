package main

import (
	"bufio"
	. "fmt"
	"github.com/Don2025/GoCode/utils"
	"os"
)

func minCostToMoveChips(position []int) int {
	var odd, even int
	for i := 0; i < len(position); i++ {
		if position[i]&1 == 0 {
			even++
		} else {
			odd++
		}
	}
	return min(odd, even)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		position := utils.StringToInts(scanner.Text())
		Printf("Output: %v\n", minCostToMoveChips(position))
	}
}
