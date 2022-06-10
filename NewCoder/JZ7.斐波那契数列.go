package main

/**
 *
 * @param n int整型
 * @return int整型
 */
func Fibonacci(n int) int {
	f, g := 0, 1
	for ; n > 0; n-- {
		g += f
		f = g - f
	}
	return f
}

/*
 * 运行时间：2ms 超过100.00%用Go提交的代码
 * 占用内存：852KB 超过40.76%用Go提交的代码
**/
