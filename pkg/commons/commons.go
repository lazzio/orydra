package commons

import (
	"regexp"
	"strings"
)

// ToSnakeCase converts a string to snake_case
func ToSnakeCase(input string) string {
	// Insère un "_" avant chaque majuscule sauf si elle est au début
	re := regexp.MustCompile(`([a-z0-9])([A-Z])`)
	snake := re.ReplaceAllString(input, `${1}_${2}`)

	// Convertit le tout en minuscules
	return strings.ToLower(snake)
}

// ConvertTabtoString converts a slice of strings to a comma-separated string
func ConvertTabtoString(tab []string) string {
	return strings.Join(tab, ",")
}
