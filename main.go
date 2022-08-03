package main

import (
	"fmt"
	"github.com/Don2025/GoCode/templates"
)

func main() {
	nums := templates.PrimeSieve(1024)
	fmt.Printf("Output: %v\n", nums)
	fmt.Printf("512: %v\n", nums[512])
}
