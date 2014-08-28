package main

/**
 * You can find more documentation how to use GoQuery in http://godoc.org/github.com/PuerkitoBio/goquery
 */

import (
	"fmt"
	"log"

	"github.com/PuerkitoBio/goquery"
	"encoding/json"
	"os"
)

type Article struct {
	Link  string
	Title string
}

func NewsScraper() {
	doc, err := goquery.NewDocument("http://korrespondent.net/")
	if err != nil {
		log.Fatal(err)
	}

	var articles = make([]Article, 0)

	doc.Find(".time-articles .article__title").Each(func(i int, s *goquery.Selection) {

		link, _ := s.Find("a").Attr("href")
		title := s.Find("a").Text()

		a := Article{link, title}
		articles = append(articles, a)
	})

	b, err := json.MarshalIndent(articles, "", "  ")
	if err != nil {
		fmt.Println("error:", err)
	}

	os.Stdout.Write(b)
}

func main() {
	NewsScraper()
}
