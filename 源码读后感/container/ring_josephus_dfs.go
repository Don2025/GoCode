package main

import "fmt"

//sum当前总人数, 数到k出圈, 最后存活人数n
func josephus(sum, k, n int) int {
	if n == 1 {
		return (sum + k - 1) % sum
	} else {
		return (josephus(sum-1, k, n-1) + k) % sum
	}
}

func main() {
	var sum, k int
	_, err := fmt.Scanf("%d %d", &sum, &k)
	if err != nil {
		return
	}
	virgin := true
	for i := 1; i <= sum; i++ {
		if virgin {
			fmt.Printf("%d", josephus(sum, k, i)+1)
			virgin = false
		} else {
			fmt.Printf(",%d", josephus(sum, k, i)+1)
		}
	}
	println()
}
