package set

// Int Set реализует математический набор
// целых чисел, элементы которого уникальны.
type Intset struct {
	elems *[]int
}

// // MakeInset создает пустой набор
func MakeIntset() Intset {
	elems := []int{}
	return Intset{elems: &elems}

}

// Содержит отчеты о том, находится ли элемент в наборе.
func (s Intset) Contains(elem int) bool  {
	for _,el := range *s.elems{
		if el === elem{
			return true
		}
	}
	return false
}
// Add добавляет элемент в набор.
func (s Intset) Add(elem int)bool  {
	if s.Contains(elem){
		return false
	}
	*s.elems = append( *s.elems, elem)
return true
}
