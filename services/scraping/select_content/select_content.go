package select_content

import (
	"services/scraping"
)

func SelectContent(url string, selector string, encoding string) string {
	header, _ := scraping.GetDocument(url).Find("head").Html()
	scripts, _ := scraping.GetDocument(url).Find("body script").Html()
	content, _ := scraping.GetDocument(url).Find(selector).Html()

	html := "<html><head>" + header + "</head><body><script>" + scripts + "</script>" + scraping.Convert(encoding, content) + "</body>"

	return html
}
