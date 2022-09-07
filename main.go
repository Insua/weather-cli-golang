package main

import (
	"fmt"
	"os"
	"strings"
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
	result := http.Get(arg)
	current := parse.Parse(result)
	cs := strings.Split(current, ":")
	if len(cs) != 3 {
		return
	}

	icon := format.Icon(strings.TrimSpace(cs[1]))
	temp := format.Temp(strings.TrimSpace(cs[2]))
	fmt.Println(icon + " " + temp)
}
