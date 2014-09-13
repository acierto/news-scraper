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

func doPost(url string, post_data io.Reader) {
	r, _ := http.NewRequest("POST", fmt.Sprintf("http://localhost:%v/%s", port, url), post_data)

	r.Header.Add("Content-Type", "application/json")
	r.Header.Add("Accept", "application/json")

	resp, _ := client.Do(r)
	defer resp.Body.Close()

	_, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("%s", err)
	}
}

func CreateDB(name string) {
	post_data := strings.NewReader("{\"name\":\"" + name + "\"}")
	doPost("_api/database", post_data)
}

func CreateCollection(dbName string, collectionName string) {
	post_data := strings.NewReader("{\"name\":\"" + collectionName + "\"}")
	doPost(fmt.Sprintf("_db/%s/_api/collection", dbName), post_data)
}

func ImportDocuments(dbName string, collectionName string, createCollection bool, body string) {
	post_data := strings.NewReader(body)
	uri :=fmt.Sprintf("_db/%s/_api/import?type=array&collection=%s&createCollection=%t", dbName, collectionName, createCollection)
	doPost(uri, post_data)
}
