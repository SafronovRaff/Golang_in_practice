package main

import (
	"fmt"
	"sort"
)

/*
Счетчик цифр
Напишите программу, которая считает, сколько раз каждая цифра встречается в числе. Гарантируется, что на вход подаются только положительные целые числа,
не выходящие за диапазон int.
Sample Input:
12823
Sample Output:
1:1 2:2 3:1 8:1
*/
func main() {
	var number int
	_, err := fmt.Scanf("%d", &number)
	if err != nil {
		fmt.Print("input error")
	}
	// Посчитайте, сколько раз каждая цифра встречается
	// в числе `number`. Запишите результат в карту `counter`,
	// где ключом является цифра числа, а значением -
	// сколько раз она встречается
	counter := make(map[int]int)
	// ...
	for number != 0 { // цикл пока не 0
		num := number % 10   // создаём переменую остатка от деления входного числа
		counter[num]++       // добавляем в мапу переменную в виде ключа и в случае повторения увеличиваем значение на +1
		number = number / 10 // убираем по одному символу с конца за цикл
	}

	printCounter(counter)
}

// printCounter печатает карту в формате
// key1:val1 key2:val2 ... keyN:valN
func printCounter(counter map[int]int) {
	digits := make([]int, 0)
	for d := range counter {
		digits = append(digits, d)
	}
	sort.Ints(digits)
	for idx, digit := range digits {
		fmt.Printf("%d:%d", digit, counter[digit])
		if idx < len(digits)-1 {
			fmt.Print(" ")
		}
	}
	fmt.Print("\n")
}
