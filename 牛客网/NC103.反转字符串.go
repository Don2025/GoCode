package main

/**
 * 反转字符串
 * @param str string字符串
 * @return string字符串
 */
func solve(str string) string {
	ans := []byte{}
	for i := len(str) - 1; i >= 0; i-- {
		ans = append(ans, str[i])
	}
	return string(ans)
}

/*
 * 运行时间：3ms 超过27.70%用Go提交的代码
 * 占用内存：1104KB 超过13.98%用Go提交的代码
**/
