package articles

import (
	"db"
	"fmt"
	"strings"
)

var dbName = "news-scraper"
var collectionName = "articles"

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
	example := fmt.Sprintf("{\"collection\": \"articles\", \"example\" :  { \"Link\" : \"%s\" }  }", link)
	return db.FindByExample(dbName, strings.NewReader(example))
}
