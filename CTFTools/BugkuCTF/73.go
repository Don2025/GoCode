package main

import (
	"fmt"
	"strconv"
	"strings"
)

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
	println(UnicodeToASCII("&#102;&#108;&#97;&#103;&#123;&#101;&#48;&#48;&#97;&#97;&#101;&#48;&#56;&#97;&#51;&#48;&#99;&#52;&#100;&#100;&#102;&#55;&#98;&#52;&#102;&#98;&#52;&#52;&#49;&#99;&#52;&#100;&#53;&#54;&#54;&#54;&#99;&#125;"))
}
