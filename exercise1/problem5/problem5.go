package problem5

import (
	"sort"
)

//products({"Computer" : 600, "TV" : 800, "Radio" : 50}, 300) ➞ ["TV","Computer"]

func products(entry map[string]int, minPrice int) []string {
	keys := make([]string, 0, len(entry))

	for key := range entry {
		if minPrice > entry[key] {
			continue
		}

		keys = append(keys, key)
	}

	sort.Slice(keys, func(i, j int) bool {
		if entry[keys[i]] == entry[keys[j]] {
			return keys[i] < keys[j]
		}

		return entry[keys[i]] > entry[keys[j]]
	})

	return keys
}
