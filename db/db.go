package db

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

func CreateDB() {
	apiUrl := "http://localhost:8529"
	resource := "/_api/database"

	u, _ := url.ParseRequestURI(apiUrl)
	u.Path = resource
	urlStr := fmt.Sprintf("%v", u)

	client := &http.Client{}
	post_data := strings.NewReader("{\"name\":\"news-scraper\"}")
	r, _ := http.NewRequest("POST", urlStr, post_data)

	r.Header.Add("Content-Type", "application/json")
	r.Header.Add("Accept", "application/json")

	resp, _ := client.Do(r)
	defer resp.Body.Close()
	fmt.Println(resp.Status)
}
