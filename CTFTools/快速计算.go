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
	client := http.Client{}
	return client.Do(request)
}

func main() {
	response, err := httpRequest("https://1254-e9566a93-6f41-45c3-be56-a1fc4678145c.do-not-trust.hacking.run")
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
