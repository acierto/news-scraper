package articles

import (
	"db"
	"fmt"
	"strings"
	"bytes"
	"model"
	"encoding/json"
)

var dbName = "news-scraper"
var collectionName = "articles"

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func CreateDB() {
	db.CreateDB(dbName)
}

func CreateCollection() {
	db.CreateCollection(dbName, collectionName)
}

func ImportDocuments(createCollection bool, body string) {
	db.ImportDocuments(dbName, collectionName, createCollection, body)
}

func FindArticle(link string) string {
	example := fmt.Sprintf("{\"collection\": \"%s\", \"example\" :  { \"Link\" : \"%s\" }  }", collectionName, link)
	return db.FindByExample(dbName, strings.NewReader(example))
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
	return db.ReadCollection(dbName, strings.NewReader(example))
}

func AddArticle(article *model.Article) {
	json, err := json.MarshalIndent(article, "", "  ")
	check(err)
	db.CreateDocument(dbName, collectionName, bytes.NewReader(json))
}
