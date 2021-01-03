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
}

func encodeText(text []rune) (string, []string) {
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

func removeString(array []string, index int) []string {
	return append(array[:index], array[index+1:]...)
}

func removeRune(array []rune, index int) []rune {
	return append(array[:index], array[index+1:]...)
}

func decodeWord(word []rune, wordLength int, encodedWords []string) (string, []string) {
	for index, encodedWord := range encodedWords {
		if len(encodedWord) != wordLength {
			continue
		}

		encoded := []rune(encodedWord)
		if word[0] != encoded[0] && word[wordLength-1] != encoded[wordLength-1] {
			continue
		}

		found := false
		partOfEncoded := encoded[1 : wordLength-1]
		for i := 1; i < wordLength-1; i++ {
			for j, letter := range partOfEncoded {
				if letter == word[i] {
					partOfEncoded = removeRune(partOfEncoded, j)
					found = true
					break
				}
			}

			if !found {
				break
			}
		}

		if found {
			return encodedWord, removeString(encodedWords, index)
		}
	}

	return string(word), encodedWords
}

func decodeText(text []rune, encodedWords []string) string {
	var currentWord []rune
	var newString []rune

	for _, item := range text {
		if unicode.IsPunct(item) || unicode.IsSpace(item) {
			currentWordLength := len(currentWord)
			if currentWordLength >= 4 {
				var decoded string
				decoded, encodedWords = decodeWord(currentWord, currentWordLength, encodedWords)
				currentWord = []rune(decoded)
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

	return string(newString)
}

func main() {
	test := `This is a long looong test sentence,
with some big (biiiiig) words!`

	encodedText, encodedWords := encodeText([]rune(test))
	fmt.Println(encodedText, encodedWords)
	fmt.Println(decodeText([]rune(encodedText), encodedWords))
}
