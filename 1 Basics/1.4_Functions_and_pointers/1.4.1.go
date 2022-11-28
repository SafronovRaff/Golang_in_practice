package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

/*
Фильтр для коллекции
Напишите функцию filter(), которая фильтрует срез целых чисел с помощью функции-предиката и возвращает отфильтрованный срез.
Функция-предикат вызывается для каждого элемента исходного среза. Если она возвращает true, элемент попадает в отфильтрованный срез. Если возвращает false — не попадает.
Считайте исходный срез из стандартного ввода с помощью готовой функции readInput(). Затем выполните на нем filter().
В качестве предиката используйте функцию, которая возвращает true только для четных чисел. Напечатайте отфильтрованный срез.
Гарантируется, что на вход подаются только целые числа.
Sample Input:
1 2 3 4 5 6
Sample Output:
[2 4 6]
*/
func filter(predicate func(int) bool, iterable []int) []int {
	// отфильтруйте `iterable` с помощью `predicate`
	// и верните отфильтрованный срез
	arr := make([]int, 0)
	for _, value := range iterable {
		if predicate(value) {
			arr = append(arr, value)
		}
	}
	return arr
}

func main() {
	src := readInput()
	// отфильтруйте `src` так, чтобы остались только четные числа
	// и запишите результат в `res`

	res := filter(s, src)
	// res := filter(func (i int) bool { return i%2 == 0}, src)

	fmt.Println(res)
}

func s(i int) bool { //
	return i%2 == 0
}

// ┌─────────────────────────────────┐
// │ не меняйте код ниже этой строки │
// └─────────────────────────────────┘

// readInput считывает целые числа из `os.Stdin`
// и возвращает в виде среза
// разделителем чисел считается пробел
func readInput() []int {
	var nums []int
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		num, _ := strconv.Atoi(scanner.Text())
		nums = append(nums, num)
	}
	return nums
}
