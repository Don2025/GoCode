package main

import (
	"bufio"
	"bytes"
	"os"
	"strings"
)

func originalDigits(s string) string {
	m := make(map[rune]int)
	for _, x := range s {
		m[x]++
	}
	cnt := [10]int{}
	cnt[0] = m['z']
	cnt[2] = m['w']
	cnt[4] = m['u']
	cnt[6] = m['x']
	cnt[8] = m['g']
	cnt[3] = m['h'] - cnt[8]
	cnt[5] = m['f'] - cnt[4]
	cnt[7] = m['s'] - cnt[6]
	cnt[1] = m['o'] - cnt[0] - cnt[2] - cnt[4]
	cnt[9] = m['i'] - cnt[5] - cnt[6] - cnt[8]
	ans := strings.Builder{}
	for i, x := range cnt {
		ans.Write(bytes.Repeat([]byte{byte('0' + i)}, x))
	}
	return ans.String()
}

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		println(originalDigits(input.Text()))
	}
}

/*
 * 执行用时：8ms 在所有Go提交中击败了75.76%的用户
 * 占用内存：4.9MB 在所有Go提交中击败了69.70%的用户
**/
