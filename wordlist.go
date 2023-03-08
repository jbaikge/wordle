package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
)

const WordLength = 5

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

	scanner := bufio.NewScanner(resp.Body)
	for scanner.Scan() {
		if len(scanner.Bytes()) == WordLength {
			fmt.Fprintf(f, "%s\n", scanner.Text())
		}
	}

	return scanner.Err()
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
		// Calling scanner.Text() twice is twice as fast as storing the word
		// in a variable
		if len(scanner.Bytes()) == WordLength {
			words = append(words, scanner.Text())
		}
	}
	return words, scanner.Err()
}
