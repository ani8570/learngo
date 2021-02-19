package tutohttp

import (
	"io/ioutil"
	"net/http"

	"github.com/ani8570/learngo/010.file/checkfunc"
)

// Gethttp : get http
func Gethttp(s string) string {
	//res, err := http.Get("https://www.google.co.kr/")
	res, err := http.NewRequest("GET", s, nil)
	checkfunc.CheckErr(err)
	//checkCode(res)
	res.Header.Add("User_Agent", "Crwaler")
	client := &http.Client{}
	resp, err := client.Do(res)
	checkfunc.CheckErr(err)
	defer resp.Body.Close()

	bytes, _ := ioutil.ReadAll(resp.Body)
	str := string(bytes)
	return str
}
