package main

/**
 *
 * @param x int整型
 * @return int整型
 */
func sqrt(x int) int {
	if x <= 1 {
		return x
	}
	l, r := 1, x
	for l <= r {
		m := l + (r-l)>>1
		if m < x/m {
			l = m + 1
		} else if m > x/m {
			r = m - 1
		} else {
			return m
		}
	}
	return r
}

/*
 * 运行时间：4ms 超过19.77%用Go提交的代码
 * 占用内存：1096KB 超过12.03%用Go提交的代码
**/
