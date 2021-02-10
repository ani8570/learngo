package main

import (
	"errors"
	"fmt"
	"time"
)

var errRequestFailed = errors.New("Request failed")

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
