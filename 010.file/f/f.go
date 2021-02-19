package f

import (
	"log"
	"net/http"
)

// Member : name and age
type Member struct {
	Name   string
	Age    int
	Active bool
}

// Members : Lots of Member
type Members struct {
	Member []Member
}

// Person : name and age
type Person struct {
	Name string
	Age  int
}

// CheckErr : err check
func CheckErr(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

// CheckCode : Code check
func CheckCode(resp *http.Response) {
	if resp.StatusCode != 200 {
		log.Fatalln("Request failed with Status:", resp.StatusCode)
	}
}
