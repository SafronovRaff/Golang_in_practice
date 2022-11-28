package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

/*
Напишите программу, которая принимает на вход фразу и составляет аббревиатуру по первым буквам слов:

Today I learned → TIL
Высшее учебное заведение → ВУЗ
Кот обладает талантом → КОТ
Если слово начинается не с буквы, игнорируйте его:
Ар 2 Ди #2 → АД
Разделителями слов считаются только пробельные символы. Дефис, дробь и прочие можно не учитывать:
Анна-Мария Волхонская → АВ
Sample Input:
Today I learned
Sample Output:
TIL
*/
func main() {
	phrase := readString()
	abbr := []rune{}
	// 1. Разбейте фразу на слова, используя `strings.Fields()`
	// 2. Возьмите первую букву каждого слова и приведите
	//    ее к верхнему регистру через `unicode.ToUpper()`
	// 3. Если слово начинается не с буквы, игнорируйте его
	//    проверяйте через `unicode.IsLetter()`
	// 4. Составьте слово из получившихся букв и запишите его
	//    в переменную `abbr`
	// ...
	phr := strings.Fields(phrase)
	for _, value := range phr {
		str := []rune(value)[0]
		if unicode.IsLetter(str) {
			abbr = append(abbr, unicode.ToUpper(str))
		}
	}

	//abbr := string(a) + string(b) + string(c)
	//fmt.Print(string(a))

	fmt.Println(string(abbr))
}

// ┌─────────────────────────────────┐
// │ не меняйте код ниже этой строки │
// └─────────────────────────────────┘

// readString читает строку из `os.Stdin` и возвращает ее
func readString() string {
	rdr := bufio.NewReader(os.Stdin)
	str, _ := rdr.ReadString('\n')
	return str
}
