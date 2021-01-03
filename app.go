package main

import (
	"fmt"
	"math/rand"
	"unicode"
)

func encodeWord(word []rune) []rune {
	wordLength := len(word)
	if wordLength < 4 {
		return word
	}

	if wordLength == 4 {
		word[1], word[2] = word[2], word[1]
		return word
	}

	toShuffle := word[1 : wordLength-1]
	toShuffleLength := wordLength - 2
	rand.Shuffle(toShuffleLength, func(i, j int) {
		toShuffle[i], toShuffle[j] = toShuffle[j], toShuffle[i]
	})

	return word
}

func processText(text []rune) string {
	var currentWord []rune
	var newString []rune

	for _, item := range text {
		if unicode.IsPunct(item) || unicode.IsSpace(item) {
			currentWord = encodeWord(currentWord)
			for _, letter := range currentWord {
				newString = append(newString, letter)
			}
			currentWord = []rune{}

			newString = append(newString, item)
			continue
		}

		currentWord = append(currentWord, item)
	}

	return string(newString)
}

func main() {
	test := `This is a long looong test sentence,
with some big (biiiiig) words!`
	fmt.Println(processText([]rune(test)))
}
