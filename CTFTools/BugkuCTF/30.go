package main

import (
	"encoding/base64"
)

func main() {
	flag, _ := base64.StdEncoding.DecodeString("ZmxhZ3swMWEyNWVhM2ZkNjM0OWM2ZTYzNWExZDAxOTZlNzVmYn0=")
	println(string(flag))
}
