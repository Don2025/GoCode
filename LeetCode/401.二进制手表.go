package main

import (
	"bufio"
	"fmt"
	"math/bits"
	"os"
	"strconv"
)

func readBinaryWatch(turnedOn int) []string {
	var ans []string
	var h, m uint8
	for h = 0; h < 12; h++ {
		for m = 0; m < 60; m++ {
			if bits.OnesCount8(h)+bits.OnesCount8(m) == turnedOn {
				ans = append(ans, fmt.Sprintf("%d:%02d", h, m))
			}
		}
	}
	return ans
}

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		n, _ := strconv.Atoi(input.Text())
		ans := readBinaryWatch(n)
		for _, x := range ans {
			fmt.Printf("%s ", x)
		}
		println()
	}
}

/*
 * 执行用时：0ms 在所有Go提交中击败了100.00%的用户
 * 占用内存：2MB 在所有Go提交中击败了75.56%的用户
**/
