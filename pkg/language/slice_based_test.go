package language_test

import (
	"github.com/google/go-cmp/cmp"
	"johnicholas.com/orkes-interview/pkg/language"
	"testing"
)

func TestSliceBased(t *testing.T) {
	tests := []struct {
		name     string
		language []string
		prefix   string
		want     []string
	}{{
		name:     "the results can be multiple",
		language: []string{"ab", "abc"},
		prefix:   "a",
		want:     []string{"ab", "abc"},
	}, {
		name:     "the results include the prefix itself, if it is a word",
		language: []string{"a", "bb", "ccc"},
		prefix:   "ccc",
		want:     []string{"ccc"},
	}, {
		name:     "no results if there are no words starting with that prefix",
		language: []string{"a", "aa", "aaa"},
		prefix:   "x",
		want:     []string{},
	}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			toTest := language.FreshSliceBased()
			for _, word := range tt.language {
				toTest.Add(word)
			}
			got := []string{}
			completions := toTest.WordsStartingWith(tt.prefix)
			for word := range completions {
				got = append(got, word)
			}
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("WordsStartingWith mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
