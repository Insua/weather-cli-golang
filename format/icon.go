package format

import (
	"github.com/gogf/gf/os/gtime"
	"github.com/gogf/gf/util/gconv"
	"github.com/thoas/go-funk"
	"strings"
)

func Icon(weather, sunrise, sunset string) string {
	if funk.ContainsString([]string{"多云", "阴", "Dreary (Overcast)"}, weather) {
		return ""
	}
	if weather == "Fog" {
		return ""
	}
	if funk.ContainsString([]string{"晴", "Intermittent Clouds", "Mostly Sunny", "Partly Sunny", "Hazy Sunshine", "Hot"}, weather) {
		sunriseArr := strings.Split(sunrise, ":")
		sunsetArr := strings.Split(sunset, ":")
		if len(sunriseArr) != 2 || len(sunsetArr) != 2 {
			return ""
		}
		sunriseHour := gconv.Uint(sunriseArr[0])
		sunriseMinute := gconv.Uint(sunriseArr[1])
		sunsetHour := gconv.Uint(sunsetArr[0])
		sunsetMinute := gconv.Uint(sunsetArr[1])
		now := gtime.Now()
		nowHour := gconv.Uint(now.Format("G"))
		nowMinute := gconv.Uint(now.Format("i"))
		if nowHour == sunsetHour && nowMinute >= sunsetMinute {
			return ""
		}
		if nowHour > sunsetHour {
			return ""
		}
		if nowHour == sunriseHour && nowMinute <= sunriseMinute {
			return ""
		}
		if nowHour < sunriseHour {
			return ""
		}
		return ""
	}
	if funk.Contains([]string{"中雪", "小雪", "阵雪"}, weather) {
		return ""
	}
	if weather == "Showers" {
		return ""
	}
	if strings.Contains(weather, "T-Storms") {
		return ""
	}
	if strings.Contains(weather, "Rain") {
		return ""
	}
	if strings.Contains(weather, "Windy") {
		return ""
	}
	if strings.Contains(weather, "Flurries") {
		return ""
	}
	if strings.Contains(weather, "Ice") {
		return ""
	}
	if strings.Contains(weather, "Sleet") {
		return ""
	}
	if strings.Contains(weather, "Cold") {
		return ""
	}
	if strings.Contains(weather, "Clear") {
		return ""
	}
	if strings.Contains(weather, "Moonlight") {
		return ""
	}
	if strings.Contains(weather, "Thunderstorms") {
		return ""
	}
	return weather
}
