package main

/**
 * You can find more documentation how to use GoQuery in http://godoc.org/github.com/PuerkitoBio/goquery
 */

import (
	"github.com/PuerkitoBio/goquery"
	"encoding/json"
	"os"
	"enc"
	"model"

	"net/http"
	"fmt"
	"io/ioutil"
	"bytes"
	"github.com/go-martini/martini"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func readInput() []model.InputElement {
	var input = make([]model.InputElement, 0)
	dat, err := ioutil.ReadFile("input.json")
	check(err)
	json.Unmarshal(dat, &input)
	return input
}

func rawElementValue(s *goquery.Selection, rules []string) string {
	for i := 0; i < len(rules); i++ {
		if rules[i] == "Find" {
			i++
			s = s.Find(rules[i])
		} else if rules[i] == "Attr" {
			i++
			link, _ := s.Attr(rules[i])
			return link
		} else if rules[i] == "Text" {
			return s.Text()
		}
	}

	return ""
}

func convert(Charset string, Text string) string {
	if Charset == "UTF-8" {
		return Text
	} else if Charset == "cp1251" {
		buffer := bytes.NewBufferString("")

		for _, char := range []byte(Text) {
			var ch = cp1251_utf8.Utf(char)
			fmt.Fprintf(buffer, "%c", ch)
		}

		return string(buffer.Bytes())
	}

	return Text
}

func findElementValue(s *goquery.Selection, rules []string, Charset string) string {
	return convert(Charset, rawElementValue(s, rules))
}

func findImage(Link string, FindImg []string) string {
	doc, err := goquery.NewDocument(Link)
	check(err)

	src := ""

	for imgNum := 0; imgNum < len(FindImg); imgNum++ {
		doc.Find(FindImg[imgNum]).Each(func(i int, s *goquery.Selection) {
			src, _ = s.Attr("src")
		})
		if src != "" {
			return src
		}
	}
	return src
}

func collectArticles(doc *goquery.Document, inputElement model.InputElement) []model.Article {
	var articles = make([]model.Article, 0)

	doc.Find(inputElement.Find).Each(func(i int, s *goquery.Selection) {
		link := findElementValue(s, inputElement.Link, inputElement.Charset)
		title := findElementValue(s, inputElement.Title, inputElement.Charset)
		img := findImage(link, inputElement.FindImg)
		a := model.Article{link, title, img}
		articles = append(articles, a)
	})
	return articles
}

func createJsonFile(jsonInput []byte) {
	f, err := os.Create("web/json/articles.json")
	check(err)

	defer f.Close()
	_, err = f.Write(jsonInput)
	check(err)
}

func NewsScraper() {

	var sourceArticles = make([]model.SourceArticle, 0)

	for _, inputElement := range readInput() {

		doc, err := goquery.NewDocument(inputElement.Source)
		check(err)

		var articles = collectArticles(doc, inputElement)
		sourceArticle := model.SourceArticle{inputElement.Source, articles}
		sourceArticles = append(sourceArticles, sourceArticle)
	}

	json, err := json.MarshalIndent(sourceArticles, "", "  ")
	check(err)

	createJsonFile(json)
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
//	NewsScraper()

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
