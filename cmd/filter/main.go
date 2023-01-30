package main

import (
	"bufio"
	"fmt"
	"io"
	"johnicholas.com/orkes-interview/pkg/language"
	"log"
	"os"
)

type Language interface {
	Add(word string)
	WordsStartingWith(prefix string) <-chan string
}

// loads the language with data from the file
func load(l Language, f *os.File) {
	bf := bufio.NewReader(f)
	for {
		line, isPrefix, err := bf.ReadLine()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		if isPrefix {
			log.Fatal("Error: Unexpected long line reading", f.Name())
		}

		l.Add(string(line))
	}
}

func main() {
	l := language.FreshSliceBased()

	wordsFile, err := os.Open("words_alpha.txt")
	if err != nil {
		log.Fatal("Error: Unexpected error opening words_alpha.txt")
	}
	load(l, wordsFile)

	f := os.Stdin // TODO: consider command line options like, "-" for stdin, otherwise os.Open(path). Cobra?

	bf := bufio.NewReader(f)
	for {
		line, isPrefix, err := bf.ReadLine()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		if isPrefix {
			log.Fatal("Error: Unexpected long line reading", f.Name())
		}

		completions := l.WordsStartingWith(string(line))
		for word := range completions {
			fmt.Println(word)
		}
	}
}
