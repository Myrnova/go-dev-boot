package CH9

import "strings"

func countDistinctWords(messages []string) int {
	wordMap := make(map[string]int)
	for _, msg := range messages {
		msg = strings.ToLower(msg)
		words := strings.Fields(msg)
		for _, word := range words {
			wordMap[word]++
		}
	}
	return len(wordMap)
}

//Distinct Words
//Complete the countDistinctWords function using a map. It should take a slice of strings and return the
//total count of distinct words across all the strings. Assume words are separated by spaces.
//Casing should not matter. (e.g., "Hello" and "hello" should be considered the same word).
//
//For example:
//
//messages := []string{"Hello world", "hello there", "General Kenobi"}
//count := countDistinctWords(messages)
//Copy icon
//count should be 5 as the distinct words are "hello", "world", "there", "general" and "kenobi" irrespective of casing.
