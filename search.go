package main

import (
	"strings"
)

func search(words []string, guesses []string) (results []string) {
	results = make([]string, len(words))
	copy(results, words)
	for i := 0; i < len(guesses); i += 2 {
		results = filter(results, guesses[i], guesses[i+1])
	}
	return
}

func filter(words []string, guess string, layout string) (filtered []string) {
	filtered = make([]string, 0, len(words))

	for _, word := range words {
		if match(word, guess, layout) {
			filtered = append(filtered, word)
		}
	}

	return
}

func match(word string, guess string, layout string) bool {
	for i, color := range layout {
		wordChar, guessChar := word[i], guess[i]

		// Greens: Must-have letters in precise positions
		if color == 'g' && wordChar != guessChar {
			return false
		}

		// Yellows: Must-have letters but not in their current positions
		if color == 'y' && wordChar == guessChar {
			return false
		}

		index := strings.IndexByte(word, guessChar)
		if color == 'y' && index == -1 {
			return false
		}

		// Blacks: Should not appear anywhere
		if color == 'b' && index != -1 {
			return false
		}
	}
	return true
}
