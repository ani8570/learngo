package main

import (
	"errors"
	"fmt"
	"time"
)

var errRequestFailed = errors.New("Request failed")

// func main() {

// 	var results = map[string]string{}

// 	urls := []string{
// 		"https://www.airbnb.co.kr/",
// 		"https://www.google.com/",
// 		"https://www.amazon.com/",
// 		"https://www.instagram.com/",
// 		"https://www.reddit.com/",
// 		"https://www.facebook.com/",
// 	}
// 	for _, url := range urls {
// 		//fmt.Println(url)
// 		result := "OK"
// 		err := hitURL(url)
// 		if err != nil {
// 			result = "FAILED"
// 		}
// 		results[url] = result
// 	}
// 	for url, result := range results {
// 		fmt.Println(url, result)
// 	}
// }

// func hitURL(url string) error {
// 	fmt.Println("Checking: ", url)
// 	resp, err := http.Get(url)
// 	if err != nil || resp.StatusCode >= 400 {
// 		fmt.Println(nil, resp.StatusCode)
// 		return errRequestFailed
// 	}
// 	return nil
// }

func main() {
	c := make(chan string)
	people := [4]string{"nico", "flynn", "ssss", "asd"}
	for _, person := range people {
		go isSexy(person, c)
	}
	for i := 0; i < len(people); i++ {
		fmt.Println(<-c)
	}

}

func sexyCount(name string) {
	for i := 0; i < 10; i++ {
		fmt.Println(name, " is sexy", i)
		time.Sleep(time.Second)
	}

}

func isSexy(person string, c chan string) {
	c <- person + " is sexy"
}
