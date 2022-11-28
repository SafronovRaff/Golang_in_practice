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

	phr := strings.Fields(phrase) // разбиваем строку на строки
	for _, value := range phr {   // ренжуем значения строк
		str := rune(value[0])      //создаём переменную(срез) первого символа строки
		if unicode.IsLetter(str) { // если символ буква то ... если нет то пропускаем
			abbr = append(abbr, unicode.ToUpper(str)) // добавляем эту букву в верхнем регистре в массив
		}
	}

	fmt.Println(string(abbr)) // конвертируем массив рун в стринг
}

// readString читает строку из `os.Stdin` и возвращает ее
func readString() string {
	rdr := bufio.NewReader(os.Stdin)
	str, _ := rdr.ReadString('\n')
	return str
}
