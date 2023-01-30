package language

import "github.com/dghubble/trie"

type trieBasedLanguage struct {
	underlying trie.Trier
}

// Constructs and returns a fresh trie-based language
func FreshTrieBased() *trieBasedLanguage {
	return &trieBasedLanguage{
		underlying: trie.NewRuneTrie(),
	}
}

func (t *trieBasedLanguage) Add(word string) {
	t.underlying.Put(word, t)
}

func (t *trieBasedLanguage) WordsStartingWith(prefix string) <-chan string {
	c := make(chan string)
	go func() {
		// This isn't quite right, because we are ignoring prefix.
		// t.WalkPath(prefix, func(key string, value interface{}) error {
		t.underlying.Walk(func(key string, value interface{}) error {
			c <- key
			return nil
		})
		close(c)
	}()
	return c
}
