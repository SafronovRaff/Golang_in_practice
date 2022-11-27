package main

import (
	"fmt"
)

/*
Укорот байтовой строки
Напишите программу, которая укорачивает строку до указанной длины и добавляет в конце многоточие:
text = Eyjafjallajokull, width = 6 → Eyjafj...
Если строка не превышает указанной длины, менять ее не следует:
text = hello, width = 6 → hello
Гарантируется, что в исходной строке text используются только однобайтовые символы без пробелов, а длина width строго больше 0.
Sample Input:
Eyjafjallajokull 6
Sample Output:

Eyjafj...
*/
func main() {
	var text string
	var width int
	_, err := fmt.Scanf("%s %d", &text, &width)
	if err != nil {
		fmt.Print("input error")
	}
	bytes := []byte(text)
	if width < len(bytes) {
		fmt.Println(string(append(bytes[:width], '.', '.', '.')))
	} else {
		fmt.Println(text)
	}

}

/*	res := text
	if len(text) > width {
	res = text[:width] + "..."
	}*/
