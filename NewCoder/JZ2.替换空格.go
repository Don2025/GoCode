package main

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 * @param s string字符串
 * @return string字符串
 */
func replaceSpace(s string) string {
	str := make([]rune, 0)
	for _, it := range s {
		if it == ' ' {
			str = append(str, '%', '2', '0')
		} else {
			str = append(str, it)
		}
	}
	return string(str)
}

/*
 * 运行时间：3ms 超过51.63%用Go提交的代码
 * 占用内存：860KB 超过79.74%用Go提交的代码
**/
