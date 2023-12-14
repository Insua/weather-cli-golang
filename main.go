package main

import (
	"fmt"
	"github.com/gogf/gf/util/gconv"
	"os"
	"strings"
	"weather/db"
	"weather/format"
	"weather/http"
	"weather/parse"

	"github.com/gogf/gf/os/gcmd"
)

func main() {
	arg := gcmd.GetArg(1)
	if len(gcmd.GetArg(1)) == 0 {
		fmt.Print("need city")
		os.Exit(0)
	}
	db.Init()
	result := http.Get(arg)
	current := parse.Parse(result)

	weather := strings.TrimSpace(current.Weather)
	fullTemp := strings.TrimSpace(current.Temp)
	icon := format.Icon(weather)
	temp := format.Temp(fullTemp)
	if len(weather) > 0 && len(fullTemp) > 0 {
		db.Record(weather, gconv.Uint(strings.Trim(fullTemp, "°C")))
	}
	fmt.Println(icon + " " + temp)
}
