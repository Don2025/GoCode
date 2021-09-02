package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	decodeBytes, _ := base64.StdEncoding.DecodeString("RGFqaURhbGlfSmlud2FuQ2hpamk=")
	flag := "flag{" + fmt.Sprintf("%s", string(decodeBytes)) + "}"
	println(flag)
}
