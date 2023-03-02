# Wordle Solver

Threw this together to quickly scan some words and give ideas on what the solution might be.

It defaults to using the standard dictionary word list (if you have one installed) located at `/usr/share/dict/words`. If the default wordlist is not available, one gets downloaded to `words.txt` in the current directory and you are instructed to use the `-words` argument.

`wordle` takes arguments in pairs with the first argument being the guessed word followed by the response Wordle returned. The response is formatted as a 5-letter string from the following letters:

- `g`: Green letters
- `y`: Yellow letters
- `b`: Black/Grey letters

## Example

`wordle later bybyb` Will return all the words with an "a" and an "e" in them, but not in their current positions.

`wordle later bybyb shade bbybg` Will filter the words from the previous guess down further to satisfy the result of the second guess
