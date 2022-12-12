package main

import (
	"fmt"
)

// Int Set реализует математический набор
// целых чисел, элементы которого уникальны.
type IntSet map[int]struct{}

// // MakeInset создает пустой набор
func MakeIntSett() IntSet {
	return IntSet{}
}

// Содержит отчеты о том, находится ли элемент в наборе.
func (s IntSet) Containss(elem int) bool {
	_, val := s[elem]
	return val
}

// Add добавляет элемент в набор.
func (s IntSet) Addd(elem int) bool {
	if s.Containss(elem) {
		return false
	}
	s[elem] = struct{}{}
	return true
}

func main() {
	sett := MakeIntSett()

	sett.Addd(5)
	fmt.Println(sett.Containss(5)) // true

	fmt.Println(sett.Containss(42))
	// false

	// элементы множества уникальны,
	// добавить дважды один и тот же элемент не получится
	added := sett.Addd(5)
	fmt.Println(added)
	// false
}
