package main

import (
	"encoding/base64"
)

func main() {
	pwd, _ := base64.StdEncoding.DecodeString("TVNEUzQ1NkFTRDEyM3p6")
	println(string(pwd))
}
