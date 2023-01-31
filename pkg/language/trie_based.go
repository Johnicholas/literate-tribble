package language

import (
	trie "github.com/tchap/go-patricia/v2/patricia"
)

type trieBasedLanguage struct {
	underlying *trie.Trie
}

// Constructs and returns a fresh trie-based language
func FreshTrieBased() *trieBasedLanguage {
	return &trieBasedLanguage{
		underlying: trie.NewTrie(),
	}
}

func (t *trieBasedLanguage) Add(word string) {
	// we (currently) do not have values associated with each word
	dontCare := 1
	t.underlying.Insert(trie.Prefix(word), dontCare)
}

func (t *trieBasedLanguage) WordsStartingWith(prefix string) <-chan string {
	c := make(chan string)
	go func() {
		// This isn't quite right, because we are ignoring prefix.
		// t.WalkPath(prefix, func(key string, value interface{}) error {
		t.underlying.VisitSubtree(trie.Prefix(prefix), func(key trie.Prefix, value trie.Item) error {
			c <- string(key)
			return nil
		})
		close(c)
	}()
	return c
}
