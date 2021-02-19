package main

import (
	"fmt"
	"os"

	"github.com/ani8570/learngo/010.file/proceed"
)

func main() {
	// a := tutohttp.Gethttp("http://csharp.tips/feed/rss")
	// fmt.Println(a)

	// p := f.Person{Name: "Lee", Age: 10}
	// b := tutohttp.Posthttp()
	// fmt.Println(b)
	// b = tutohttp.Posthttp2()
	// fmt.Println(b)
	// b = tutohttp.Postjson(p)
	// fmt.Println(b)
	// b = tutohttp.PostXML(p)
	// fmt.Println(b)
	// b = tutohttp.Postdetail(p)
	// fmt.Println(b)

	// m := f.Member{Name: "Lee", Age: 10, Active: true}
	// c := proceed.Enjson(m)
	// fmt.Println(c)
	// //if you want watch json data, it can convert to string
	// fmt.Println(string(c))
	// d := proceed.Dejson(c)
	// fmt.Println(d)
	path, _ := os.Getwd()
	s := path + "\\proceed\\test.xml"
	fmt.Println(s)
	e := proceed.ReadXML(s)
	fmt.Println(e)
}
