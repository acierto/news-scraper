package articles

import (
	"db"
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
