package wordcount

import (
	"regexp"
	"strings"
)

// count occurrence  each word
func CountMeats(text string) map[string]int {
	re := regexp.MustCompile(`\b[a-zA-Z\-]+\b`)
	words := re.FindAllString(text, -1)

	counts := make(map[string]int)
	for _, word := range words {
		word = strings.ToLower(word)
		counts[word]++
	}

	return counts

}
