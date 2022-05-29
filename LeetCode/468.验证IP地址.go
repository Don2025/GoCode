package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func validIPAddress(queryIP string) string {
	if arr := strings.Split(queryIP, "."); len(arr) == 4 {
		for _, ip := range arr {
			if len(ip) > 1 && ip[0] == '0' {
				return "Neither"
			}
			if ipNum, err := strconv.Atoi(ip); ipNum > 255 || err != nil {
				return "Neither"
			}
		}
		return "IPv4"
	}
	if arr := strings.Split(queryIP, ":"); len(arr) == 8 {
		for _, ip := range arr {
			if len(ip) < 1 || len(ip) > 4 {
				return "Neither"
			}
			if _, err := strconv.ParseUint(ip, 16, 64); err != nil {
				return "Neither"
			}
		}
		return "IPv6"
	}
	return "Neither"
}

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		println(validIPAddress(input.Text()))
	}
}

/*
 * 执行用时：0ms 在所有Go提交中击败了100.00%的用户
 * 占用内存：1.9MB 在所有Go提交中击败了81.17%的用户
**/
