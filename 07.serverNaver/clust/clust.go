package clust

import (
	"encoding/csv"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/djimenez/iconv-go"
)

// News information
type News struct {
	newsURL string
	title   string
}

// Clustnews naver news by num
func Clustnews(term string) {
	var baseURL string = "https://news.naver.com/main/list.nhn?mode=LSD&mid=sec&sid1=10" + term
	var URLs []News
	URLs = getPages(URLs, baseURL)
	writeNews(URLs, term)
}

func getPages(URLs []News, baseURL string) []News {

	res, err := http.Get(baseURL)
	checkErr(err)
	checkCode(res)
	defer res.Body.Close()
	doc, err := goquery.NewDocumentFromReader(res.Body)
	checkErr(err)

	searchCards := doc.Find("dt.photo a")
	searchCards.Each(func(i int, s *goquery.Selection) {

		f1, _ := s.Attr("href")
		ss := s.Find("img")
		f2, _ := ss.Attr("alt")
		f2, _ = iconv.ConvertString(string(f2), "euc-kr", "utf-8")
		ClearString(&f2)
		fmt.Println(f1, f2)
		//URLs = append([]string{res}, URLs...)
		URLs = append(URLs, News{newsURL: f1, title: f2})
	})
	return URLs
}
func writeNews(URLs []News, term string) {
	newHead := map[string]string{"0": "정치", "1": "경제", "2": "사회", "3": "생활and문화", "4": "세계", "5": "ITand과학"}
	file, err := os.Create("Naver_News_" + newHead[term] + ".csv")
	checkErr(err)

	w := csv.NewWriter(file)
	defer w.Flush() //함수가 끝나는 시점에서 실행됨

	headers := []string{"Part", "URL", "Title"}

	errW := w.Write(headers)
	checkErr(errW)

	for idx, url := range URLs {
		urlSlice := []string{strconv.Itoa(idx + 1), url.newsURL, url.title}
		erruW := w.Write(urlSlice)
		checkErr(erruW)
	}
}

//ClearString make clean string
func ClearString(str *string) {
	s := strings.TrimSpace(*str)
	s = strings.Replace(s, "\"", "", -1)
	*str = s
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
