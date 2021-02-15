package main

import (
	"learngo/07.serverNaver/clust"
	"os"

	"github.com/labstack/echo/v4"
)

func handleHome(c echo.Context) error {
	return c.File("home.html")
}
func handleclust(c echo.Context) error {
	term := c.FormValue("term")
	newHead := map[string]string{"0": "정치", "1": "경제", "2": "사회", "3": "생활and문화", "4": "세계", "5": "ITand과학"}
	clust.ClearString(&term)
	defer os.Remove("Naver_News_" + newHead[term] + ".csv")
	if term >= "1" && term <= "5" {
		clust.Clustnews(term)
		return c.Attachment("Naver_News_"+newHead[term]+".csv", "Naver_News_"+newHead[term]+".csv")
	}
	return c.File("home.html")
}

func main() {
	e := echo.New()
	e.GET("/", handleHome)
	e.POST("/clust", handleclust)
	e.Logger.Fatal(e.Start(":8080"))

}
