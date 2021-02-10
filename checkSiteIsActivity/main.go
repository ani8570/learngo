package main

import (
	"errors"
	"fmt"
	"net/http"
)

type requestResult struct {
	url    string
	status string
}

var errRequestFailed = errors.New("Request failed")

func main() {

	var results = map[string]string{}

	c := make(chan requestResult)
	urls := []string{
		"https://www.airbnb.co.kr/",
		"https://www.google.com/",
		"https://www.amazon.com/",
		"https://www.instagram.com/",
		"https://www.reddit.com/",
		"https://www.facebook.com/",
		"https://www.a.com",
	}
	for _, url := range urls {
		go hitURL(url, c)
	}
	for i := 0; i < len(urls); i++ {
		result := <-c
		results[result.url] = result.status
	}
	for key, value := range results {
		fmt.Println(key, value)
	}
}

func hitURL(url string, c chan requestResult) {
	resp, err := http.Get(url)
	status := "OK"
	if err != nil || resp.StatusCode >= 400 {
		status = "Fail"
	}
	c <- requestResult{url: url, status: status}
}
