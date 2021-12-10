package main

/**
 * 给出一个有序的整数数组 A 和有序的整数数组 B ，请将数组 B 合并到数组 A 中，变成一个有序的升序数组。
 * 1.保证 A 数组有足够的空间存放 B 数组的元素， A 和 B 中初始的元素数目分别为 m 和 n，A的数组空间大小为 m+n
 * 2.不要返回合并的数组，将数组 B 的数据合并到 A 里面就好了
 * 3. A 数组在[0,m-1]的范围也是有序的
**/

/**
 *
 * @param A int整型一维数组
 * @param B int整型一维数组
 * @return void
 */
func merge(A []int, m int, B []int, n int) {
	index := m + n - 1
	for m > 0 && n > 0 {
		if A[m-1] >= B[n-1] {
			A[index] = A[m-1]
			m--
		} else {
			A[index] = B[n-1]
			n--
		}
		index--
	}
	for n > 0 {
		A[index] = B[n-1]
		n--
		index--
	}
}

/*
 * 运行时间：4ms 超过76.98%用Go提交的代码
 * 占用内存：1004KB 超过83.17%用Go提交的代码
**/
