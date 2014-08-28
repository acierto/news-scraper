package main

/**
 * You can find more documentation how to use GoQuery in http://godoc.org/github.com/PuerkitoBio/goquery
 */

import (
	"github.com/PuerkitoBio/goquery"
	"encoding/json"
	"os"

	"io/ioutil"
)

type InputElement struct {
	Source string
	Find   string
	Link   []string
	Title  []string
}

type Article struct {
	Link  string
	Title string
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func readInput() []InputElement {
	var input = make([]InputElement, 0)
	dat, err := ioutil.ReadFile("input.json")
	check(err)
	json.Unmarshal(dat, &input)
	return input
}

func findElementValue(s *goquery.Selection, rules []string) string {

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

func NewsScraper() {

	var articles = make([]Article, 0)

	for _, inputElement := range readInput() {

		doc, err := goquery.NewDocument(inputElement.Source)
		check(err)

		doc.Find(inputElement.Find).Each(func(i int, s *goquery.Selection) {
			link := findElementValue(s, inputElement.Link)
			title := findElementValue(s, inputElement.Title)

			a := Article{link, title}
			articles = append(articles, a)
		})
	}

	json, err := json.MarshalIndent(articles, "", "  ")
	check(err)

	os.Stdout.Write(json)
}

func main() {
	NewsScraper()
}
