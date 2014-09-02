package main

/**
 * You can find more documentation how to use GoQuery in http://godoc.org/github.com/PuerkitoBio/goquery
 */

import (
	"net/http"
	"services/scraping"
	"services/liveupdate"
	"github.com/go-martini/martini"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	scraping.Scrape()dd
	liveupdate.CronLatestNews()

	m := martini.Classic()
	m.Use(martini.Static("web"))

	m.Get("/read-html", func(req *http.Request) string {
			urlValues := req.URL.Query()["url"]
			if len(urlValues) > 0 {
				return scraping.SelectContent(urlValues[0], ".leftCell.mainContentBlock")
			}
			return ""
		})

	m.Run()
}
