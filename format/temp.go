package format

import "strings"

func Temp(temp string) string {
	return strings.ReplaceAll(temp, "C", "°C")
}
