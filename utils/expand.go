package utils

import "strings"

func ExpandTemplate(str string, prefix string, region string, message string) string {
	replacer := strings.NewReplacer("%p", prefix, "%r", region, "%m", message)

	return replacer.Replace(str)
}
