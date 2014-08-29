package main

/**
 * You can find more documentation how to use GoQuery in http://godoc.org/github.com/PuerkitoBio/goquery
 */

import (
	"github.com/PuerkitoBio/goquery"
	"encoding/json"
	"os"
	"cp1251_utf8"

	"fmt"
	"io/ioutil"
	"bytes"
)

type InputElement struct {
	Source  string
	Find    string
	Link    []string
	Title   []string
	Charset string
}

func (self *InputElement) Initialize() { if self.Charset == "" { self.Charset = "UTF-8" } }

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

func NewsScraper() {

	var articles = make([]Article, 0)

	for _, inputElement := range readInput() {

		doc, err := goquery.NewDocument(inputElement.Source)
		check(err)
		fmt.Println(doc.Url)

		doc.Find(inputElement.Find).Each(func(i int, s *goquery.Selection) {
			link := findElementValue(s, inputElement.Link, inputElement.Charset)
			title := findElementValue(s, inputElement.Title, inputElement.Charset)

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
