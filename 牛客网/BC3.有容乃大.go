package main

import (
	"fmt"
	"unsafe"
)

func main() {
	fmt.Printf("The size of short is %d bytes.\n", unsafe.Sizeof(int16(1)))
	fmt.Printf("The size of int is %d bytes.\n", unsafe.Sizeof(int32(1)))
	fmt.Printf("The size of long is %d bytes.\n", unsafe.Sizeof(int64(1)))
	fmt.Printf("The size of long long is %d bytes.\n", unsafe.Sizeof(int(1)))
}
/*
 * 运行时间：3ms 超过10.00%用Go提交的代码
 * 占用内存：800KB 超过55.00%用Go提交的代码
**/