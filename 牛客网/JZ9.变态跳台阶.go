package main

func jumpFloorII(number int) int {
	return 1 << (number - 1)
}

/*
 * 运行时间：5ms 超过3.32%用Go提交的代码
 * 占用内存：940KB 超过9.97%用Go提交的代码
**/
