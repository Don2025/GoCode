package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func httpRequest(url string) (*http.Response, error) {
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	request.Header.Add("Referer", "https://www.Sycsecret.com")
	request.Header.Add("User-Agent", "Syclover")
	request.Header.Add("X-Forwarded-For", "127.0.0.1")
	client := http.Client{}
	return client.Do(request)
}

func main() {
	response, err := httpRequest("http://node4.buuoj.cn:25280/Secret.php")
	if err != nil {
		fmt.Printf("http get error: %s", err)
		return
	}
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Printf("read error: %s", err)
	}
	fmt.Println(string(body))
}
