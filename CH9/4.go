package CH9

import (
	"golang.org/x/exp/utf8string"
)

func getNameCounts(names []string) map[rune]map[string]int {
	nameCounts := make(map[rune]map[string]int)
	for _, name := range names {
		s := utf8string.NewString(name)
		r := s.At(0)
		if _, ok := nameCounts[r]; !ok {
			nameCounts[r] = make(map[string]int)
		}
		nameCounts[r][name]++
	}
	return nameCounts
}
