package main

import (
	"fmt"
)

/*
Язык по коду
Напишите программу, которая определяет название языка по его коду. Правила:
en → English
fr → French
ru или rus → Russian
иначе → Unknown
Sample Input:
en
Sample Output:
English
*/
func main() {
	var code string
	_, err := fmt.Scan(&code)
	if err != nil {
		fmt.Print("input error")
	}
	// определите полное название языка по его коду
	// и запишите его в переменную `lang`
	// ...
	var lang string
	switch code {
	case "en":
		lang = "English"
	case "fr":
		lang = " French"
	case "ru", "rus":
		lang = "Russian"

	default:
		lang = "Unknown"
	}

	fmt.Println(lang)
}
