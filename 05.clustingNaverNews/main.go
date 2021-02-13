package main

import (
	"encoding/csv"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

var baseURL string = "https://news.naver.com/main/list.nhn?mode=LS2D&mid=shm&sid1=101&sid2=259"

func main() {
	var URLs []string
	URLs = getPages(URLs)
	//fmt.Println(URLs)
	writeNews(URLs)
}

func getPages(URLs []string) []string {

	res, err := http.Get(baseURL)
	checkErr(err)
	checkCode(res)
	defer res.Body.Close()
	doc, err := goquery.NewDocumentFromReader(res.Body)
	checkErr(err)

	searchCards := doc.Find("dl dt a")
	searchCards.Each(func(i int, s *goquery.Selection) {
		res, _ := s.Attr("href")
		//URLs = append([]string{res}, URLs...)
		URLs = append(URLs, res)
	})
	return URLs
}
func writeNews(URLs []string) {
	file, err := os.Create("Naver_finace.csv")
	checkErr(err)

	w := csv.NewWriter(file)
	defer w.Flush() //함수가 끝나는 시점에서 실행됨

	headers := []string{"Part", "URL"}

	errW := w.Write(headers)
	checkErr(errW)

	for idx, url := range URLs {
		urlSlice := []string{strconv.Itoa(idx + 1), url}
		erruW := w.Write(urlSlice)
		checkErr(erruW)

	}
}

func clearString(str string) string {
	return strings.Join(strings.Fields(strings.TrimSpace(str)), " ")
}
func checkErr(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func checkCode(resp *http.Response) {
	if resp.StatusCode != 200 {
		log.Fatalln("Request failed with Status:", resp.StatusCode)
	}
}
