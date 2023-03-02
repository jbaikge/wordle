package main

import "testing"

func TestMatch(t *testing.T) {
	tests := []struct {
		Guess   string
		Word    string
		Layout  string
		Matches bool
	}{
		{
			Guess:   "later",
			Word:    "shade",
			Layout:  "bybyb",
			Matches: true,
		},
		{
			Guess:   "later",
			Word:    "alter",
			Layout:  "bybyb",
			Matches: false,
		},
		{
			Guess:   "later",
			Word:    "alter",
			Layout:  "yyggg",
			Matches: true,
		},
	}

	for _, test := range tests {
		if match(test.Word, test.Guess, test.Layout) != test.Matches {
			t.Fatalf("match(%s, %s, %s) != %v", test.Word, test.Guess, test.Layout, test.Matches)
		}
	}
}
