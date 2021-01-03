package main

import (
	"fmt"
)

type EncodedText struct {
	text         string
	encodedWords []string
}

func (e EncodedText) String() string {
	return fmt.Sprintf("\n---weird---\n%s\n---weird---\n%v", e.text, e.encodedWords)
}
