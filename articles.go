package main

import (
	"fmt"
	"strings"
	"bytes"
	"encoding/json"
)

var dbName = "news-scraper"
var collectionName = "articles"

func ImportArticles(createCollection bool, body string) {
	ImportDocuments(dbName, collectionName, createCollection, body)
}

func FindArticle(link string) string {
	example := fmt.Sprintf("{\"collection\": \"%s\", \"example\" :  { \"Link\" : \"%s\" }  }", collectionName, link)
	return FindByExample(dbName, strings.NewReader(example))
}

func HasArticle(link string) bool {
	articleJson := FindArticle(link)

	var objmap map[string]*json.RawMessage
	err := json.Unmarshal([]byte(articleJson), &objmap)
	check(err)

	var articles = make([]interface{}, 0)
	err = json.Unmarshal(*objmap["result"], &articles)
	check(err)

	return len(articles) > 0
}

func GetAllArticles() string {
	example := fmt.Sprintf("{\"collection\": \"%s\"}", collectionName)
	return ReadCollection(dbName, strings.NewReader(example))
}

func AddArticle(article *Article) {
	json, err := json.MarshalIndent(article, "", "  ")
	check(err)
	CreateDocument(dbName, collectionName, bytes.NewReader(json))
}
