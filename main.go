package main

/**
 * You can find more documentation how to use GoQuery in http://godoc.org/github.com/PuerkitoBio/goquery
 */

import (
	"net/http"
	"github.com/go-martini/martini"
	"encoding/json"
	"log"
)

func check(e error) {
	if e != nil {
		log.Printf(e.Error())
	}
}

func routing() {
	m := martini.Classic()
	m.Use(martini.Static("web"))

	m.Get("/read-articles", func(req *http.Request) string {
			return GetAllArticles()
		})

	m.Get("/read-html", func(req *http.Request) string {
			urlValues := req.URL.Query()["url"]
			selectors := req.URL.Query()["selector"]
			charset := req.URL.Query()["charset"]
			if len(urlValues) > 0 && len(selectors) > 0 {
				return SelectContent(urlValues[0], selectors[0], charset[0])
			}
			return ""
		})

	m.Run()
}

func importScrapedArticles() {
	articles := Scrape()

	json, err := json.MarshalIndent(articles, "", "  ")
	check(err)
	ImportArticles(true, string(json))
}

func prepareDB() {
	if !ExistsDB() {
		CreateDB()
		CreateCollection(articlesCollectionName)
	}
}

func runJob() {
	prepareDB()
	importScrapedArticles()
	CronLatestNews()
}

func main() {
	runJob()
	routing()
}
