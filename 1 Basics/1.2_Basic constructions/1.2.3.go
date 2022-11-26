package main

import (
	"fmt"
	"strings"
)

/*
строки
Программа принимает на вход строку source и число times. Требуется склеить source саму с собой times раз и вернуть результат:
source = x, times = 3 → xxx
source = omm, times = 2 → ommomm
Sample Input:
a 5
Sample Output:
aaaaa
*/

func main() {
	var source string
	var times int
	// гарантируется, что значения корректные
	_, err := fmt.Scan(&source, &times)
	if err != nil {
		fmt.Print("input error")
	}
	// возьмите строку `source` и повторите ее `times` раз
	// запишите результат в `result`
	// ...
	result := strings.Repeat(source, times)
	fmt.Println(result)
}

/*var result string
for i := 0; i < times; i++ {
result += source
}*/
