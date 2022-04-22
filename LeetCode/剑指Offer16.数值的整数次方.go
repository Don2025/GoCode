package main

func myPow(x float64, n int) float64 {
	var quickMul func(float64, int) float64
	quickMul = func(x float64, n int) float64 {
		ans := 1.0
		for n > 0 {
			if n&1 == 1 {
				ans *= x
			}
			x *= x
			n >>= 1
		}
		return ans
	}
	if n >= 0 {
		return quickMul(x, n)
	}
	return 1 / quickMul(x, -n)
}

/*
 * 执行用时：0ms 在所有Go提交中击败了100.00%的用户
 * 占用内存：1.9MB 在所有Go提交中击败了66.67%的用户
**/
