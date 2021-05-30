package LeetCode

func isPowerOfTwo(n int) bool {
	//利用负数在计算机中的存储方式，检测一个数是否为2的幂只要判断 n&(-n) 是否等于 n 本身即可
	return n > 0 && (n&(-n)) == n
}
