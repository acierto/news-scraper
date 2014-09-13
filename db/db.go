package db

import (
	"io"
	"fmt"
	"net/http"
	"strings"
	"io/ioutil"
)

var port = 8529
var client = &http.Client{}

func send(method string, url string, data io.Reader) string {
	r, _ := http.NewRequest(method, fmt.Sprintf("http://localhost:%v/%s", port, url), data)

	r.Header.Add("Content-Type", "application/json")
	r.Header.Add("Accept", "application/json")

	resp, _ := client.Do(r)
	defer resp.Body.Close()

	contents, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("%s", err)
	}

	return string(contents)
}

func doPost(url string, post_data io.Reader) {
	send("POST", url, post_data)
}

func FindByExample(dbName string, example io.Reader) string {
	return send("PUT", fmt.Sprintf("_db/%s/_api/simple/by-example", dbName), example)
}

func CreateDB(name string) {
	post_data := strings.NewReader("{\"name\":\"" + name + "\"}")
	doPost("_api/database", post_data)
}

func CreateCollection(dbName string, collectionName string) {
	post_data := strings.NewReader("{\"name\":\"" + collectionName + "\"}")
	doPost(fmt.Sprintf("_db/%s/_api/collection", dbName), post_data)
}

func ReadCollection(dbName string, data io.Reader) string {
	return send("PUT", fmt.Sprintf("_db/%s/_api/simple/all", dbName), data)
}

func CreateDocument(dbName string, collectionName string, data io.Reader) {
	send("POST", fmt.Sprintf("_db/%s/_api/document?collection=%s", dbName, collectionName), data)
}

func ImportDocuments(dbName string, collectionName string, createCollection bool, body string) {
	post_data := strings.NewReader(body)
	uri :=fmt.Sprintf("_db/%s/_api/import?type=array&collection=%s&createCollection=%t", dbName, collectionName, createCollection)
	doPost(uri, post_data)
}
