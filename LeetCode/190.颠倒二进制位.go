package main

func reverseBits(num uint32) uint32 {
	var ans uint32
	for i := 32; i > 0; i-- {
		ans <<= 1
		ans += num & 1
		num >>= 1
	}
	return ans
}

/*
 * 执行用时：0ms 在所有Go提交中击败了100.00%的用户
 * 占用内存：2.4MB 在所有Go提交中击败了76.98%的用户
**/
