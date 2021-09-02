package main

import "encoding/hex"

// hexToStr converts Hex code to text string.
func hexToStr(num string) string {
	str, _ := hex.DecodeString(num)
	return string(str)
}

// strToHex converts text string to Hex code.
func strToHex(text string) string {
	return hex.EncodeToString([]byte(text))
}

func main() {
	println(hexToStr("73646E6973635F32303138"))
	println(strToHex("sdnisc_2018")) //Just test
}
