package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

type Color string

type Letter struct {
	Char  string
	Color Color
}

type Word [5]Letter

const (
	ColorReset  = Color("\033[0m")
	ColorGreen  = Color("\033[32m")
	ColorYellow = Color("\033[33m")
	ColorGrey   = Color("\033[90m")
)

var wordsPath = "/usr/share/dict/words"

func (w *Word) Set(word string) {
	for i := range w {
		w[i].Char = word[i : i+1]
		w[i].Color = ColorGrey
	}
}

func (w Word) String() (s string) {
	for _, letter := range w {
		s += string(letter.Color) + letter.Char + string(ColorReset)
	}
	return
}

func init() {
	flag.StringVar(&wordsPath, "path", wordsPath, "Path to words file")
	flag.Parse()
}

func isLowercase(word string) bool {
	for i := range word {
		if word[i] < 'a' || word[i] > 'z' {
			return false
		}
	}
	return true
}

func match(word, green, yellow, grey string) (w Word, matched bool) {
	if len(word) != 5 {
		return
	}

	if !isLowercase(word) {
		return
	}

	w.Set(word)

	// Handle green letters
	for i := range green {
		if green[i] == '.' {
			continue
		}
		if word[i] != green[i] {
			return
		}
		w[i].Color = ColorGreen
	}

	for i := range yellow {
		idx := strings.IndexByte(word, yellow[i])
		if idx == -1 {
			return
		}
		w[idx].Color = ColorYellow
	}

	matched = true

	return
}

func search(green, yellow, grey string) (results []Word, err error) {
	f, err := os.Open(wordsPath)
	if err != nil {
		return
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		word, matched := match(scanner.Text(), green, yellow, grey)
		if !matched {
			continue
		}
		results = append(results, word)
	}
	if err = scanner.Err(); err != nil {
		return
	}

	return
}

func main() {
	if flag.NArg() != 3 {
		usage("Incorrect number of arguments")
	}

	green, yellow, grey := flag.Arg(0), flag.Arg(1), flag.Arg(2)
	if len(green) != 5 {
		usage("Green argument must be 5 characters long")
	}

	results, err := search(green, yellow, grey)
	if err != nil {
		usage("Error during processing: " + err.Error())
	}

	for _, result := range results {
		fmt.Printf("%+v\n", result)
	}
}

func usage(msg string) {
	fmt.Println(msg + "\n")
	fmt.Println("Usage: wordle ..t.. la er")
	fmt.Println("")
	fmt.Println("Green letters with dots for \"any\"")
	fmt.Println("Yellow letters to include")
	fmt.Println("Grey letters to exclude")
	fmt.Println("")
	flag.Usage()
	os.Exit(1)
}
