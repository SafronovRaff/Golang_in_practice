package join

import (
	"sort"
	"strings"
)

//  Join Words объединяет слова из двух строк, удаляет дубликаты
// и возвращает результирующие слова в отсортированном порядке.

func JoinWords(first, second string) []string {
	words1 := split(first)
	words2 := split(second)
	words := join(words1, words2)
	words = lower(words)
	words = sorte(words)
	words = uniq(words)
	return words
}

// split разбивает строку на слова.
func split(str string) []string {
	return strings.Fields(str)
}

// соединение соединяет два фрагмента слова в один фрагмент.
func join(words1, words2 []string) []string {
	words := make([]string, len(words1)+len(words2))
	idx := 0
	for _, word := range words1 {
		words[idx] = word
		idx++
	}
	for _, word := range words2 {
		words[idx] = word
		idx++
	}
	return words
}

// lower преобразует все слова в нижний регистр.

func lower(words []string) []string {
	for idx, word := range words {
		words[idx] = strings.ToLower(word)
	}
	return words
}

// sorted сортирует все слова в алфавитном порядке.
func sorte(words []string) []string {
	sort.Strings(words)
	return words
}

// uniq удаляет повторяющиеся слова.
func uniq(words []string) []string {
	uniq := []string{}
	current := " "
	for _, word := range words {
		if word == current {
			continue
		}
		if current != "" {
			uniq = append(uniq, current)
			current = ""
		}
		current = word
	}
	if current != "" {
		uniq = append(uniq, current)
	}
	return uniq

}
