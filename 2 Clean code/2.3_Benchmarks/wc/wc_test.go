package wc

import (
	"fmt"
	"math/rand"
	"strings"
	"testing"
)

func BenchmarkWordCountRegexp(b *testing.B) {
	for _, lenght := range []int{10, 100, 1000, 10000} {

		rand.Seed(0)
		phrase := randomPhrase(lenght)
		name := fmt.Sprintf("Regexp- %d", lenght)
		b.Run(name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				WordCountRegexp(phrase)
			}
		})
	}

}

func randomPhrase(n int) string {
	words := make([]string, n)
	for i := range words {
		words[i] = randomWords(3)
	}
	return strings.Join(words, " ")
}

func randomWords(n int) string {
	const letters = "abcdefghijklmnopqrstuvwxyz"
	chars := make([]byte, n)
	for i := range chars {
		chars[i] = letters[rand.Intn(len(letters))]
	}
	return string(chars)

}
