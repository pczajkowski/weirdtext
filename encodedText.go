package weirdtext

import (
	"fmt"
	"strings"
)

type EncodedText struct {
	Text         string
	EncodedWords []string
}

const (
	prefix = "\n---weird---\n"
)

func (e EncodedText) String() string {
	return fmt.Sprintf("%s%s%s%v", prefix, e.Text, prefix, strings.Join(e.EncodedWords, " "))
}

func (e *EncodedText) FromString(serialized string) error {
	if !strings.HasPrefix(serialized, prefix) {
		return fmt.Errorf("Invalid prefix: %s", serialized)
	}

	parts := strings.Split(serialized, prefix)
	if len(parts) != 3 {
		return fmt.Errorf("Invalid string: %s", serialized)
	}

	if parts[1] == "" {
		return fmt.Errorf("There's no encoded text: %s", serialized)
	}
	e.Text = parts[1]

	if parts[2] == "" {
		return fmt.Errorf("There are no encoded words: %s", serialized)
	}
	e.EncodedWords = strings.Split(parts[2], " ")

	return nil
}
