package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"strings"
)

const Columns = 80

var wordsPath = "/usr/share/dict/words"

func init() {
	flag.StringVar(&wordsPath, "words", wordsPath, "Path to wordlist")
}

func usage(msg string) {
	fmt.Println(msg + "\n")
	fmt.Println("Usage: wordle guess1 layout1 [guess2 layout2 [...]]")
	fmt.Println("")
	fmt.Println("Provide each guess (up to 5 guesses) and the response layout")
	fmt.Println("from Wordle in the form of a string consisting of g, y, and b")
	fmt.Println("g: Green letters")
	fmt.Println("y: Yellow letters")
	fmt.Println("b: Black letters")
	fmt.Println("")
	fmt.Println("Example:")
	fmt.Println("wordle later bbbyb spine ybbbg")
	fmt.Println("")
	fmt.Println("* Meant for use with hard mode activated in Wordle")
	fmt.Println("")
	flag.Usage()
	os.Exit(1)
}

func main() {
	flag.Parse()
	if flag.NArg() < 2 {
		usage("Need at least two arguments")
	}

	if flag.NArg()%2 != 0 {
		usage("Must have an even number of arguments")
	}

	words, err := wordlist(wordsPath)
	if err != nil {
		fmt.Printf("Error processing wordlist: %v\n", err)

		if !errors.Is(err, os.ErrNotExist) {
			os.Exit(1)
		}

		fmt.Println("Downloading wordlist")
		if err = downloadWordlist(); err != nil {
			fmt.Printf("Error downloading wordlist: %v\n", err)
			os.Exit(1)
		}
		fmt.Println("Downloaded wordlist, re-run with -words words.txt")
		os.Exit(1)
	}

	results := search(words, flag.Args())

	// Pretty print wordlist
	wordsPerLine := Columns / (WordLength + 1)
	for i := 0; i < len(results); i += wordsPerLine {
		fmt.Println(strings.Join(results[i:i+wordsPerLine], " "))
	}
}
