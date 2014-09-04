package scraping

import (
	"github.com/PuerkitoBio/goquery"
	"model"
	"encoding/json"
	"os"

	"enc/cp1251_utf8"
	"fmt"
	"io/ioutil"
	"bytes"
	"time"
	"strings"
	"math"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func readInput() []model.InputElement {
	var input = make([]model.InputElement, 0)
	dat, err := ioutil.ReadFile("scrapingRules.json")
	check(err)
	json.Unmarshal(dat, &input)
	return input
}

func rawElementValue(s *goquery.Selection, rules []string) string {
	for i := 0; i < len(rules); i++ {
		if rules[i] == "Find" {
			i++
			s = s.Find(rules[i])
		} else if rules[i] == "Prev" {
			s = s.Prev()
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

func Convert(Charset string, Text string) string {
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

func findAndCovertElementValue(s *goquery.Selection, rules []string, Charset string) string {
	return Convert(Charset, rawElementValue(s, rules))
}

func findElementValue(s *goquery.Selection, rules []string) string {
	return rawElementValue(s, rules)
}

func getArticleTime(s *goquery.Selection, articleTimes *[]time.Time, inputElement model.InputElement) time.Time {
	t := findElementValue(s, inputElement.Time)

	articleTime, _ := time.Parse("2006/01/02 15:04", time.Now().UTC().Format("2006/01/02")+" "+t)

	timeZoneOffset := time.Duration(inputElement.TimeZone) * time.Hour

	if math.Abs(float64(inputElement.TimeZone)) > float64(0) {
		if articleTime.Hour() + inputElement.TimeZone <= 0 {
			articleTime = articleTime.AddDate(0, 0, 1)
		}
	}


	// time zone
	articleTime = articleTime.Add(timeZoneOffset)

	if len(*articleTimes) == 0 {
		if articleTime.After(time.Now()) {
			articleTime = articleTime.AddDate(0, 0, -1)
		}
	} else {
		var lastArticleTime = (*articleTimes)[len(*articleTimes) - 1]
		if articleTime.After(lastArticleTime) {
			articleTime = articleTime.AddDate(0, 0, -1)
		}
	}

	if (len(*articleTimes) > 2) {
		(*articleTimes) = (*articleTimes)[1:]
	}

	(*articleTimes) = append((*articleTimes), articleTime)

	return articleTime
}

func getLink(s *goquery.Selection, inputElement model.InputElement) string {
	link := findElementValue(s, inputElement.Link)
	if strings.HasPrefix(link, "/") {
		link = inputElement.Source + link
	}
	return link
}

func collectArticles(doc *goquery.Document, inputElement model.InputElement) []model.Article {
	var articles = make([]model.Article, 0)

	var articleTimes = make([]time.Time, 0)
	doc.Find(inputElement.Find).Each(func(i int, s *goquery.Selection) {
		title := findAndCovertElementValue(s, inputElement.Title, inputElement.Charset)
		link := getLink(s, inputElement)
		articleTime := getArticleTime(s, &articleTimes, inputElement)

		a := model.Article{inputElement.Source, inputElement.ContentSelector, link, title, articleTime, inputElement.Charset}
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

func GetDocument(htmlLink string) *goquery.Document {
	doc, err := goquery.NewDocument(htmlLink)
	check(err)
	return doc
}

func Scrape() {
	var articles = make([]model.Article, 0)

	for _, inputElement := range readInput() {
		var doc = GetDocument(inputElement.Source)
		var collectedArticles = collectArticles(doc, inputElement)
		articles = append(articles, collectedArticles...)
	}

	json, err := json.MarshalIndent(articles, "", "  ")
	check(err)

	createJsonFile(json)
	fmt.Println("Scraping is successfully finished at:" + time.Now().String())
}
