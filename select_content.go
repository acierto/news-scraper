package main

func SelectContent(url string, selector string, encoding string) string {
	header, _ := GetDocument(url).Find("head").Html()
	scripts, _ := GetDocument(url).Find("body script").Html()
	content, _ := GetDocument(url).Find(selector).Html()

	html := "<html><head>" + header + "</head><body><script>" + scripts + "</script>" + Convert(encoding, content) + "</body>"

	return html
}
