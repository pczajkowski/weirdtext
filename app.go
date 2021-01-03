package main

import (
	"fmt"
	"math/rand"
	"sort"
	"unicode"
)

func encodeWord(word []rune, wordLength int) {
	if wordLength == 4 {
		word[1], word[2] = word[2], word[1]
		return
	}

	toShuffle := word[1 : wordLength-1]
	toShuffleLength := wordLength - 2
	rand.Shuffle(toShuffleLength, func(i, j int) {
		toShuffle[i], toShuffle[j] = toShuffle[j], toShuffle[i]
	})

	return
}

func processText(text []rune) (string, []string) {
	var currentWord []rune
	var newString []rune
	var encodedWords []string

	for _, item := range text {
		if unicode.IsPunct(item) || unicode.IsSpace(item) {
			currentWordLength := len(currentWord)
			if currentWordLength >= 4 {
				beforeEncoding := string(currentWord)
				encodeWord(currentWord, currentWordLength)

				if string(currentWord) != beforeEncoding {
					encodedWords = append(encodedWords, beforeEncoding)
				}
			}

			for _, letter := range currentWord {
				newString = append(newString, letter)
			}
			currentWord = []rune{}

			newString = append(newString, item)
			continue
		}

		currentWord = append(currentWord, item)
	}

	sort.Strings(encodedWords)
	return string(newString), encodedWords
}

func main() {
	test := `This is a long looong test sentence,
with some big (biiiiig) words!`
	fmt.Println(processText([]rune(test)))
}
