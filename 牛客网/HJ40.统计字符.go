package main

import (
    "bufio"
    "fmt"
    "os"
    "unicode"
)

func main() {
    input := bufio.NewScanner(os.Stdin)
    for input.Scan() {
        s := input.Text()
        var alphaCnt, spaceCnt, digitCnt, otherCnt int
        for _, x := range s {
            if unicode.IsDigit(rune(x)) {
                digitCnt++
            } else if unicode.IsLetter(rune(x)) {
                alphaCnt++
            } else if unicode.IsSpace(rune(x)) {
                spaceCnt++
            } else {
                otherCnt++
            }
        }
        fmt.Println(alphaCnt)
        fmt.Println(spaceCnt)
        fmt.Println(digitCnt)
        fmt.Println(otherCnt)
    }
}

/*
 * 运行时间：3ms 超过51.35%用Go提交的代码
 * 占用内存：924KB 超过50.00%用Go提交的代码
**/
