package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func superPow(a int, b []int) int {
	const mod = 1337
	var exp int
	for i := 0; i < len(b); i++ {
		exp = (exp*10 + b[i]) % phi(mod)
	}
	return quickPow(a%mod, exp, mod)
}

func quickPow(x, n, mod int) int {
	if n == 0 {
		return 1
	}
	ans := quickPow(x, n/2, mod)
	ans *= ans
	ans %= mod
	if (n & 1) != 0 {
		ans *= x
		ans %= mod
	}
	return ans
}

// phi 求欧拉函数φ(n), phi就是φ的读音
func phi(n int) int {
	ans := n
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			ans = ans * (i - 1) / i
			for n%i == 0 {
				n /= i
			}
		}
	}
	if n > 1 {
		ans = ans * (n - 1) / n
	}
	return ans
}

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		a, _ := strconv.Atoi(input.Text())
		input.Scan()
		b := stringArrayToIntArray(strings.Fields(input.Text()))
		println(superPow(a, b))
	}
}

func stringArrayToIntArray(strings []string) []int {
	var arr []int
	for _, x := range strings {
		n, _ := strconv.Atoi(x)
		arr = append(arr, n)
	}
	return arr
}

/*
 * 执行用时：16ms 在所有Go提交中击败了44.87%的用户
 * 占用内存：3.6MB 在所有Go提交中击败了94.87%的用户
**/
