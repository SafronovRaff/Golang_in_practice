package main

import (
	"fmt"
	"strings"
	"sync"
	"unicode"
)

// счетчик хранит количество цифр в каждом слове.
// каждый ключ - это слово, а значение - это количество цифр в слове.
type counter map[string]int

// // подсчитывать цифры в словах подсчитывает цифры в словосочетаниях
func countDigitsInWords(phrase string) counter {
	words := strings.Fields(phrase)
	syncStats := sync.Map{}

	var wg sync.WaitGroup
	// 	wg.Add(len(words))
	// начало решения

	for _, word := range words {
		wg.Add(1) // счетчик в самом цикле
		go func(word string) {
			defer wg.Done()
			count := countDigits(word)
			syncStats.Store(word, count)
		}(word)

		wg.Wait()
	}
	// Посчитайте количество цифр в словах,
	// используя отдельную горутину для каждого слова.

	// Чтобы записать результаты подсчета,
	// используйте syncStats.Store(word, count)

	// В результате syncStats должна содержать слова
	// и количество цифр в каждом.

	// конец решения

	return asStats(syncStats)
}

// // count Digits возвращает количество цифр в строке
func countDigits(str string) int {
	count := 0
	for _, char := range str {
		if unicode.IsDigit(char) {
			count++
		}
	}
	return count
}

// // как указано, преобразует статистику из синхронизации. Сопоставить с обычной картой
func asStats(m sync.Map) counter {
	stats := counter{}
	m.Range(func(word, count any) bool {
		stats[word.(string)] = count.(int)
		return true
	})
	return stats
}

// printStats печатает слова и количество их цифр
func printStats(stats counter) {
	for word, count := range stats {
		fmt.Printf("%s: %d\n", word, count)
	}
}

func main() {
	phrase := "0ne 1wo thr33 4068"
	counts := countDigitsInWords(phrase)
	printStats(counts)
}
