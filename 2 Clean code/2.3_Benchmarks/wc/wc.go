package wc

import (
	"regexp"
	"strings"
)

// Счетчик сопоставляет слова с их количеством

type Counter map[string]int

var splitter *regexp.Regexp = regexp.MustCompile(" ")

// // Регулярное выражение Word Count подсчитывает абсолютные частоты слов в строке.
// Использует Regexp.Split() для разделения строки на слова.

func WordCountRegexp(s string) Counter {
	counter := make(Counter)
	for _, word := range splitter.Split(s, -1) {
		word = strings.ToLower(word)
		counter[word]++
	}
	return counter
}

func WordCountFields(s string) Counter {
	counter := make(Counter)
	for _, word := range strings.Fields(s) {
		word = strings.ToLower(word)
		counter[word]++
	}
	return counter
}
