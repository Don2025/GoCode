package main

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 * @param numbers int整型一维数组
 * @return int整型
 */
func duplicate(numbers []int) int {
	a := make([]bool, len(numbers))
	for _, x := range numbers {
		if !a[x] {
			a[x] = true
		} else {
			return x
		}
	}
	return -1
}

/*
 * 运行时间：7ms 超过46.09%用Go提交的代码
 * 占用内存：2152KB 超过80.00%用Go提交的代码
**/
