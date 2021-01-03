package main

import (
	"fmt"
	"strings"
)

type EncodedText struct {
	Text         string
	EncodedWords []string
}

func (e EncodedText) String() string {
	return fmt.Sprintf("\n---weird---\n%s\n---weird---\n%v", e.Text, strings.Join(e.EncodedWords, " "))
}

func (e *EncodedText) FromString(serialized string) error {
	if !strings.HasPrefix(serialized, "\n---weird---\n") {
		return fmt.Errorf("Invalid prefix: %s", serialized)
	}

	parts := strings.Split(serialized, "\n---weird---\n")
	if len(parts) != 3 {
		return fmt.Errorf("Invalid string: %s", serialized)
	}

	e.Text = parts[1]
	e.EncodedWords = strings.Split(parts[2], " ")
	return nil
}
