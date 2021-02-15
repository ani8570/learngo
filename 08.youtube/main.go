package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

var baseURL string = "https://www.youtube.com/c/rudbeckia7/videos"

func main() {
	var URLs []string
	URLs = getPages(URLs)
	//writeNews(URLs)
}

func getPages(URLs []string) []string {

	res, err := http.Get(baseURL)
	checkErr(err)
	checkCode(res)
	defer res.Body.Close()
	doc, err := goquery.NewDocumentFromReader(res.Body)
	fmt.Println(doc.Html())
	checkErr(err)
	searchCards := doc.Find("text")
	searchCards.Each(func(i int, s *goquery.Selection) {
		fmt.Println(s.Html())
		fmt.Println(s.Text())
		//res, _ := s.Attr("href")
		///html/body/ytd-app/div/ytd-page-manager/ytd-browse/ytd-two-column-browse-results-renderer/div[1]/ytd-section-list-renderer/div[2]/ytd-item-section-renderer/div[3]/ytd-grid-renderer/div[1]/ytd-grid-video-renderer['+str(idx)+']/div[1]/div[1]/div[1]/h3/a
		//URLs = append([]string{res}, URLs...)
		//URLs = append(URLs, res)
	})
	return URLs
}
func writeNews(URLs []string) {
	file, err := os.Create("youtube_rudbeckia7.csv")
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
