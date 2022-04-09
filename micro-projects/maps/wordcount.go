package main

import (
	"strings"
)

func WordCount(s string) map[string]int {
	wordList := strings.Fields(s)
	wordMap := make(map[string]int)
	for _, word := range wordList {
		wordMap[word] += 1
	}

	return wordMap
}
