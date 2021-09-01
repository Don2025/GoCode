package main

import (
	"encoding/base64"
	"log"
)

func base64Encode(text string) string {
	return base64.StdEncoding.EncodeToString([]byte(text))
}

func base64Decode(text string) string {
	decodeBytes, err := base64.StdEncoding.DecodeString(text)
	if err != nil {
		log.Fatalln(err)
	}
	return string(decodeBytes)
}

func main() {
	text := []byte("e6Z9i~]8R~U~QHE{RnY{QXg~QnQ{^XVlRXlp^XI5Q6Q6SKY8jUAA")
	//最后的AA让人联想到BASE64末尾的==
	for i := 0; i < len(text); i++ {
		text[i] -= 4 //ASCII码中=是61而A是65，盲猜一波凯撒和base64的混合加密
	}
	println(string(text))                                                         //凯撒密码解码后的字符串
	println(base64Decode("a2V5ezY4NzQzMDAwNjUwMTczMjMwZTRhNThlZTE1M2M2OGU4fQ==")) //base64解码后的字符串
}
