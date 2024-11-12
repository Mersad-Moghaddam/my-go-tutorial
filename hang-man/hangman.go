package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

var words = []string{
	"golang",
	"hangman",
	"programming",
	"development",
	"interface",
	"function",
	"variable",
	"syntax",
}

func main() {
	rand.Seed(time.Now().UnixNano())
	word := words[rand.Intn(len(words))]
	wordRunes := []rune(word)
	guessedRunes := make([]rune, len(wordRunes))
	incorrectGuesses := 0
	const maxIncorrectGuesses = 6
	guessedLetters := make(map[rune]bool)

	fmt.Println("Welcome to Hangman!")
	fmt.Println("  -----\n  |   |\n  |   O\n  |  /|\\\n  |  / \\\n  |\n  |")

	fmt.Println("Guess the word:")

	for incorrectGuesses < maxIncorrectGuesses {
		displayWord(wordRunes, guessedRunes)
		fmt.Printf("Incorrect guesses: %d/%d\n", incorrectGuesses, maxIncorrectGuesses)
		fmt.Print("Enter a letter: ")

		var guess string
		fmt.Scanln(&guess)

		if len(guess) != 1 {
			fmt.Println("Please enter a single letter.")
			continue
		}

		letter := rune(guess[0])
		if guessedLetters[letter] {
			fmt.Println("You already guessed that letter.")
			continue
		}

		guessedLetters[letter] = true

		if strings.ContainsRune(word, letter) {
			for i, r := range wordRunes {
				if r == letter {
					guessedRunes[i] = r
				}
			}
			fmt.Println("Good guess!")
		} else {
			incorrectGuesses++
			fmt.Println("Incorrect guess.")
		}

		if string(guessedRunes) == word {
			fmt.Printf("Congratulations! You've guessed the word: %s\n", word)
			return
		}
	}

	fmt.Printf("Game over! The word was: %s\n", word)
}

func displayWord(wordRunes []rune, guessedRunes []rune) {
	for i, r := range wordRunes {
		if guessedRunes[i] == 0 {
			fmt.Print("_ ")
		} else {
			fmt.Printf("%c ", r)
		}
	}
	fmt.Println()
}
