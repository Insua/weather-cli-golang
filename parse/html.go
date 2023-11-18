package parse

import (
	"github.com/PuerkitoBio/goquery"
	"strings"
)

func Parse(result []byte) string {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(string(result)))
	if err != nil {
		return ""
	}

	return doc.Find(".head .head-right p").First().Text()
}
