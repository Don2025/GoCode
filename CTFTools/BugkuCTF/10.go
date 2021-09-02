package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	decodeBytes, _ := base64.StdEncoding.DecodeString("eTB1IEFyZSBhIGhAY2tlciE=")
	flag := "f1@g{" + fmt.Sprintf("%s", string(decodeBytes)) + "}"
	println(flag)
}
