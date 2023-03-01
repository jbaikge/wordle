package main

import (
	"bufio"
	"io"
	"net/http"
	"os"
)

func downloadWordlist() (err error) {
	resp, err := http.Get("https://github.com/dwyl/english-words/raw/master/words_alpha.txt")
	if err != nil {
		return
	}
	defer resp.Body.Close()

	f, err := os.Create("words.txt")
	if err != nil {
		return
	}
	defer f.Close()

	if _, err = io.Copy(f, resp.Body); err != nil {
		return
	}
	return
}

func wordlist(path string) (words []string, err error) {
	f, err := os.Open(path)
	if err != nil {
		return
	}
	defer f.Close()

	words = make([]string, 0, 20000)
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		if word := scanner.Text(); len(word) == WordLength {
			words = append(words, scanner.Text())
		}
	}
	return words, scanner.Err()
}
