package main

import (
	"fmt"
	"strings"
	"bytes"
	"encoding/json"
)

var articlesCollectionName = "articles"

func ImportArticles(createCollection bool, body string) {
	ImportDocuments(articlesCollectionName, createCollection, body)
}

func FindArticle(link string) string {
	example := fmt.Sprintf("{\"collection\": \"%s\", \"example\" :  { \"Link\" : \"%s\" }  }", articlesCollectionName, link)
	return FindByExample(strings.NewReader(example))
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
	example := fmt.Sprintf("{\"collection\": \"%s\"}", articlesCollectionName)
	return ReadCollection(strings.NewReader(example))
}

func AddArticle(article *Article) {
	json, err := json.MarshalIndent(article, "", "  ")
	check(err)
	CreateDocument(articlesCollectionName, bytes.NewReader(json))
}
