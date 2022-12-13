package main

// не удаляйте импорты, они используются при проверке
import (
	"fmt"
	"strings"
)

/*type Words struct {
	str map[string]int
}

func MakeWords(s string) Words {
	words := strings.Fields(s)
	size := len(words) / 2
	if size > 10000 {
		size = 1000
	}
	c := make(map[string]int, size)
	for i, word := range words {
		if _, ok := c[word]; !ok {
			c[word] = i
		}
	}
	return Words{c}
}

func (w Words) Index(word string) int {

	idx, item := w.str[word]
	if item {
		return idx
	}

	return -1
}*/

func main() {
	s := "go is awesome, php is not"
	w := MakeWords(s)

	fmt.Println(w.Index("go"))
	// 0
	fmt.Println(w.Index("is"))
	// 1
	fmt.Println(w.Index("is not"))
	// -1
	fmt.Println(w.Index("python"))
	// -1
}

type Words struct {
	wmap map[string]int
}

func MakeWords(s string) Words {
	wmap := map[string]int{}
	for idx, word := range strings.Fields(s) {
		if _, found := wmap[word]; found {
			continue
		}
		wmap[word] = idx
	}
	return Words{wmap}
}

func (w Words) Index(word string) int {
	idx, found := w.wmap[word]
	if !found {
		return -1
	}
	return idx
}
