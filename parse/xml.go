package parse

import (
	"github.com/gogf/gf/encoding/gjson"

	"github.com/gogf/gf/encoding/gxml"
)

func Parse(result []byte) string {
	dr, err := gxml.ToJson(result)
	if err != nil {
		return ""
	}

	gj := gjson.New(dr)
	return gj.GetString("rss.channel.item.0.title")
}
