package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
)

func main() {
    input := bufio.NewScanner(os.Stdin)
    for input.Scan() {
        s := input.Text()
        hash := make([]int, 127)
        for _, x := range s {
            hash[x]++
        }
        for i, x := range hash {
            fmt.Print(strings.Repeat(string(rune(i)), x))
        }
        fmt.Println()
    }
}

/*
 * 运行时间：5ms 超过48.00%用Go提交的代码
 * 占用内存：968KB 超过47.00%用Go提交的代码
**/
