package main

import (
	"io"
	"fmt"
	"net/http"
	"strings"
	"io/ioutil"
	"os"
	"encoding/json"
)

var dbName = "news-scraper"
var port = 8529
var client = &http.Client{}

func send(method string, url string, data io.Reader) string {
	r, _ := http.NewRequest(method, fmt.Sprintf("http://localhost:%v/%s", port, url), data)

	r.Header.Add("Content-Type", "application/json")
	r.Header.Add("Accept", "application/json")

	resp, _ := client.Do(r)

	if resp == nil {
		fmt.Println("Cannot connect to ArangoDB. Database URL is ", fmt.Sprintf("http://localhost:%d/", port))
		os.Exit(1)
	}
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

func FindByExample(example io.Reader) string {
	return send("PUT", fmt.Sprintf("_db/%s/_api/simple/by-example", dbName), example)
}

func ExistsDB() bool {
	res := send("POST", fmt.Sprintf("_db/%s/_api/version", dbName), nil)

	var dat map[string]interface{}
	err := json.Unmarshal([]byte(res), &dat)
	check(err)

	return dat["server"] != nil
}

func CreateDB() {
	post_data := strings.NewReader("{\"name\":\"" + dbName + "\"}")
	doPost("_api/database", post_data)
}

func CreateCollection(collectionName string) {
	post_data := strings.NewReader("{\"name\":\"" + collectionName + "\"}")
	doPost(fmt.Sprintf("_db/%s/_api/collection", dbName), post_data)
}

func ReadCollection(data io.Reader) string {
	return send("PUT", fmt.Sprintf("_db/%s/_api/simple/all", dbName), data)
}

func CreateDocument(collectionName string, data io.Reader) {
	send("POST", fmt.Sprintf("_db/%s/_api/document?collection=%s", dbName, collectionName), data)
}

func ImportDocuments(collectionName string, createCollection bool, body string) {
	post_data := strings.NewReader(body)
	uri :=fmt.Sprintf("_db/%s/_api/import?type=array&collection=%s&createCollection=%t", dbName, collectionName, createCollection)
	doPost(uri, post_data)
}
