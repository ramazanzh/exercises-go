package problem4

import (
	"strings"
)

func mapping(letters []string) map[string]string {
	res := make(map[string]string, len(letters))

	for _, letter := range letters {
		res[letter] = strings.ToUpper(letter)
	}

	return res
}
