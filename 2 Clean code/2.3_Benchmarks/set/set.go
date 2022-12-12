package set

// Int Set реализует математический набор
// целых чисел, элементы которого уникальны.
type IntSet struct {
	elems *[]int
}

// // MakeInset создает пустой набор
func MakeIntSet() IntSet {
	elems := []int{}
	return IntSet{&elems}
}

// Содержит отчеты о том, находится ли элемент в наборе.
func (s IntSet) Contains(elem int) bool {
	for _, el := range *s.elems {
		if el == elem {
			return true
		}
	}
	return false
}

// Add добавляет элемент в набор.
func (s IntSet) Add(elem int) bool {
	if s.Contains(elem) {
		return false
	}
	*s.elems = append(*s.elems, elem)
	return true
}

/*func main() {
	set := MakeIntset()

	set.Add(5)
	fmt.Println(set.Contains(5)) // true

	fmt.Println(set.Contains(42))
	// false

	// элементы множества уникальны,
	// добавить дважды один и тот же элемент не получится
	added := set.Add(5)
	fmt.Println(added)
	// false
}
*/
