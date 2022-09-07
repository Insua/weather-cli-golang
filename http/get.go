package http

import (
	"time"

	"github.com/gogf/gf/frame/g"
)

func Get(city string) []byte {
	r, err := g.Client().Timeout(60 * time.Second).Get("https://rss.accuweather.com/rss/liveweather_rss.asp?metric=1&locCode=" + city)
	if err != nil {
		return []byte{}
	}
	return r.ReadAll()
}
