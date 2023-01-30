package language

import "strings"

type sliceBasedLanguage struct {
	underlying []string
}

// Constructs and returns a fresh slice based language
func FreshSliceBased() *sliceBasedLanguage {
	return &sliceBasedLanguage{
		underlying: []string{},
	}
}

func (t *sliceBasedLanguage) Add(word string) {
	t.underlying = append(t.underlying, word)
}

func (t *sliceBasedLanguage) WordsStartingWith(prefix string) <-chan string {
	c := make(chan string)
	go func() {
		for _, word := range t.underlying {
			if strings.HasPrefix(word, prefix) {
				c <- word
			}
		}
		close(c)
	}()
	return c
}
