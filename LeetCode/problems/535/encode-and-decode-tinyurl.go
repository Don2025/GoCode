package main

import (
	"strconv"
	"strings"
)

// https://leetcode.cn/problems/encode-and-decode-tinyurl/
//------------------------Leetcode Problem 535------------------------

type Codec struct {
	DataBase map[int]string
	Id       int
}

func Constructor() Codec {
	return Codec{map[int]string{}, 0}
}

// Encodes a URL to a shortened URL.
func (c *Codec) encode(longUrl string) string {
	c.Id++
	c.DataBase[c.Id] = longUrl
	return "http://tinyurl.com/" + strconv.Itoa(c.Id)
}

// Decodes a shortened URL to its original URL.
func (c *Codec) decode(shortUrl string) string {
	i := strings.LastIndexByte(shortUrl, '/')
	id, _ := strconv.Atoi(shortUrl[i+1:])
	return c.DataBase[id]
}

/**
 * Your Codec object will be instantiated and called as such:
 * obj := Constructor();
 * url := obj.encode(longUrl);
 * ans := obj.decode(url);
 */
//------------------------Leetcode Problem 535------------------------
/*
 * https://leetcode.cn/problems/encode-and-decode-tinyurl/
 * 执行用时：0ms 在所有Go提交中击败了100.00%的用户
 * 占用内存：2.4MB 在所有Go提交中击败了54.05%的用户
**/
