package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// ChineseToUnicode converts Chinese to Unicode
func ChineseToUnicode(text string) string {
	ans := strconv.QuoteToASCII(text)
	return ans[1 : len(ans)-1] //Remove quotation marks
}

// UnicodeToChinese converts Unicode to Chinese
func UnicodeToChinese(text string) string {
	arr := strings.Split(text, "\\u")
	var ans string
	for _, x := range arr {
		if len(x) < 1 {
			continue
		}
		ch, _ := strconv.ParseInt(x, 16, 32)
		ans += fmt.Sprintf("%c", ch)
	}
	return ans
}

// ASCIIToUnicode converts ASCII to Unicode
func ASCIIToUnicode(text string) string {
	var ans string
	for _, x := range text {
		ans += "&#" + strconv.Itoa(int(x)) + ";"
	}
	return ans
}

// UnicodeToASCII converts Unicode to ASCII
func UnicodeToASCII(text string) string {
	//Remove the last semicolon, then cut it into a string array according to the semicolon
	arr := strings.Split(strings.Trim(text, ";"), ";")
	var ans string
	for _, x := range arr {
		n, _ := strconv.Atoi(strings.Trim(x, "&#"))
		ans += fmt.Sprintf("%c", n)
	}
	return ans
}

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		ans := UnicodeToASCII(input.Text())
		println(ans)
	}
}
