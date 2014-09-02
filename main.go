package main

/**
 * You can find more documentation how to use GoQuery in http://godoc.org/github.com/PuerkitoBio/goquery
 */

import (
	"net/http"
	"io/ioutil"
	"services/scraping"
	"github.com/go-martini/martini"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}


func readHtmlPage(Url string) string {
	r, err := http.Get(Url)
	check(err)

	var b []byte
	b, err = ioutil.ReadAll(r.Body)
	check(err)

	var body = string(b)
	r.Body.Close()
	return body
}

func main() {
	scraping.Scrape()

	m := martini.Classic()
	m.Use(martini.Static("web"))

	m.Get("/read-html", func(req *http.Request) string {
		urlValues := req.URL.Query()["url"]
		if len(urlValues) > 0 {
			return readHtmlPage(urlValues[0])
		}
		return ""
	})

	m.Run()
}
