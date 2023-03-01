package main

const WordLength = 5

const (
	ColorReset  = Color("\033[0m")
	ColorGreen  = Color("\033[32m")
	ColorYellow = Color("\033[33m")
	ColorGrey   = Color("\033[90m")
)

type Color string

type Letter struct {
	Char  rune
	Color Color
}

type Word [WordLength]Letter

func NewWord(word string) (w Word) {
	w.Set(word)
	return
}

func (w *Word) Set(word string) {
	for i, ch := range word {
		w[i].Char = ch
		w[i].Color = ColorGrey
	}
}

func (w Word) String() (s string) {
	r := make([]rune, 0, (len(ColorReset)+1+len(ColorGreen))*WordLength)
	reset := []rune(ColorReset)
	for _, letter := range w {
		r = append(r, []rune(letter.Color)...)
		r = append(r, letter.Char)
		r = append(r, reset...)
	}
	return string(r)
}
