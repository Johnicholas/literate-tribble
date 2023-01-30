# README

The problem is, given a language, that is, a set of words such as dwyl's english-words/words_alpha.txt,
and a prefix, find the words from the language with that prefix.

To see how it works:

```bash
go build ./cmd/filter
cat prefixes.txt
./filter <prefixes.txt
```
