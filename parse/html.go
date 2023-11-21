package parse

import (
	"github.com/PuerkitoBio/goquery"
	"strings"
)

type Weather struct {
	Weather string
	Temp    string
}

func Parse(result []byte) Weather {
	w := Weather{}

	doc, err := goquery.NewDocumentFromReader(strings.NewReader(string(result)))
	if err != nil {
		return w
	}

	if ws, err := doc.Find(".real .real_weather .weather p").Last().Html(); err == nil {
		wss := strings.Split(ws, "<br/>")
		if len(wss) > 0 {
			w.Weather = wss[0]
		}
	}

	temp := doc.Find(".real .real_weather .weather p").First().Text()
	tr := make([]rune, 0)
	for _, v := range []rune(temp) {
		if v != 8451 {
			tr = append(tr, v)
		}
	}
	w.Temp = string(tr) + "°C"

	return w
}
