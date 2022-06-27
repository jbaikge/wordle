# Wordle Solver

Threw this together to quickly scan some words and give ideas on what the solution might be. It defaults to using the standard dictionary word list (if you have one installed) located at `/usr/share/dict/words`.

It takes 3 arguments: Green letters (and position), yellow letters, grey letters

For example:

`wordle ..t.. er al`

Dots space out where the `t` is located, this string must always be 5 characters.

`er` will be included in any found words

`al` will be excluded from all words
