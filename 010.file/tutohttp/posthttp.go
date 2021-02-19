package tutohttp

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/ani8570/learngo/010.file/checkfunc"
)

var postURL string = "http://httpbin.org/post"

// Posthttp : Post http
func Posthttp() string {
	reqbody := bytes.NewBufferString(("Post plain text"))
	resp, err := http.Post(postURL, "text/plain", reqbody)
	checkfunc.CheckErr(err)

	respBody, err := ioutil.ReadAll(resp.Body)
	checkfunc.CheckErr(err)
	str := string(respBody)
	return str
}

// Posthttp2 : Post http second method
func Posthttp2() string {
	resp, err := http.PostForm(postURL, url.Values{"Name": {"Lee"}, "Age": {"10"}})
	checkfunc.CheckErr(err)
	defer resp.Body.Close()
	respBody, err := ioutil.ReadAll(resp.Body)
	checkfunc.CheckErr(err)
	return string(respBody)
}

// Postjson : json form data post
func Postjson(p checkfunc.Person) string {
	pbytes, err := json.Marshal(p)
	checkfunc.CheckErr(err)
	buff := bytes.NewBuffer(pbytes)
	resp, err := http.Post(postURL, "application/json", buff)
	checkfunc.CheckErr(err)
	defer resp.Body.Close()
	respBody, err := ioutil.ReadAll(resp.Body)
	checkfunc.CheckErr(err)
	return string(respBody)
}

// PostXML : XML form data post(similar to json)
func PostXML(p checkfunc.Person) string {
	pbytes, err := json.Marshal(p)
	checkfunc.CheckErr(err)
	buff := bytes.NewBuffer(pbytes)
	resp, err := http.Post(postURL, "application/xml", buff)
	checkfunc.CheckErr(err)
	defer resp.Body.Close()
	respBody, err := ioutil.ReadAll(resp.Body)
	checkfunc.CheckErr(err)
	return string(respBody)
}

//Postdetail : Posting in more detail
func Postdetail(p checkfunc.Person) string {
	pbytes, err := json.Marshal(p)
	checkfunc.CheckErr(err)
	buff := bytes.NewBuffer(pbytes)
	req, err := http.NewRequest("POST", postURL, buff)
	checkfunc.CheckErr(err)
	req.Header.Add("Content-type", "application/xml")
	clint := &http.Client{}
	resp, err := clint.Do(req)
	checkfunc.CheckErr(err)
	defer resp.Body.Close()
	respBody, err := ioutil.ReadAll(resp.Body)
	checkfunc.CheckErr(err)
	return string(respBody)
}
