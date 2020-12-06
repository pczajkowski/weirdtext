package main

import (
	"fmt"
	"math/rand"
)

func main() {
	test := "Poświadczam"
	runes := []rune(test)
	part := runes[1 : len(runes)-1]
	partLength := len(part)
	fmt.Println(string(part))

	rand.Shuffle(partLength, func(i, j int) {
		part[i], part[j] = part[j], part[i]
	})

	fmt.Println(string(part))
}
