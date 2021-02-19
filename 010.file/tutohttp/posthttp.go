package tutohttp

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/ani8570/learngo/010.file/f"
)

var postURL string = "http://httpbin.org/post"

// Posthttp : Post http
func Posthttp() string {
	reqbody := bytes.NewBufferString(("Post plain text"))
	resp, err := http.Post(postURL, "text/plain", reqbody)
	f.CheckErr(err)

	respBody, err := ioutil.ReadAll(resp.Body)
	f.CheckErr(err)
	str := string(respBody)
	return str
}

// Posthttp2 : Post http second method
func Posthttp2() string {
	resp, err := http.PostForm(postURL, url.Values{"Name": {"Lee"}, "Age": {"10"}})
	f.CheckErr(err)
	defer resp.Body.Close()
	respBody, err := ioutil.ReadAll(resp.Body)
	f.CheckErr(err)
	return string(respBody)
}

// Postjson : json form data post
func Postjson(p f.Person) string {
	pbytes, err := json.Marshal(p)
	f.CheckErr(err)
	buff := bytes.NewBuffer(pbytes)
	resp, err := http.Post(postURL, "application/json", buff)
	f.CheckErr(err)
	defer resp.Body.Close()
	respBody, err := ioutil.ReadAll(resp.Body)
	f.CheckErr(err)
	return string(respBody)
}

// PostXML : XML form data post(similar to json)
func PostXML(p f.Person) string {
	pbytes, err := json.Marshal(p)
	f.CheckErr(err)
	buff := bytes.NewBuffer(pbytes)
	resp, err := http.Post(postURL, "application/xml", buff)
	f.CheckErr(err)
	defer resp.Body.Close()
	respBody, err := ioutil.ReadAll(resp.Body)
	f.CheckErr(err)
	return string(respBody)
}

//Postdetail : Posting in more detail
func Postdetail(p f.Person) string {
	pbytes, err := json.Marshal(p)
	f.CheckErr(err)
	buff := bytes.NewBuffer(pbytes)
	req, err := http.NewRequest("POST", postURL, buff)
	f.CheckErr(err)
	req.Header.Add("Content-type", "application/xml")
	clint := &http.Client{}
	resp, err := clint.Do(req)
	f.CheckErr(err)
	defer resp.Body.Close()
	respBody, err := ioutil.ReadAll(resp.Body)
	f.CheckErr(err)
	return string(respBody)
}
