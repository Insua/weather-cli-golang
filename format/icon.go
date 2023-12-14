package format

import (
	"strings"

	"github.com/thoas/go-funk"
)

func Icon(weather string) string {
	if funk.ContainsString([]string{"多云", "阴", "Dreary (Overcast)"}, weather) {
		return ""
	}
	if weather == "Fog" {
		return ""
	}
	if funk.ContainsString([]string{"晴", "Intermittent Clouds", "Mostly Sunny", "Partly Sunny", "Hazy Sunshine", "Hot"}, weather) {
		return ""
	}
	if funk.Contains([]string{"中雪", "小雪"}, weather) {
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
