package problem2

import (
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func capitalize(names []string) []string {
	var result []string

	for _, name := range names {
		result = append(result, cases.Title(language.English, cases.Compact).String(name))
	}

	return result
}
