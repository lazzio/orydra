package commons

import (
	"regexp"
	"strings"
)

func ToInterface(input string) interface{} {
	return input
}

func ToSnakeCase(input string) string {
	// Insère un "_" avant chaque majuscule sauf si elle est au début
	re := regexp.MustCompile(`([a-z0-9])([A-Z])`)
	snake := re.ReplaceAllString(input, `${1}_${2}`)

	// Convertit le tout en minuscules
	return strings.ToLower(snake)
}

func ConvertTabtoString(tab []string) string {
	return strings.Join(tab, ",")
}
