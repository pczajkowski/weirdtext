package main

import (
	"testing"
)

func TestDecodeEncode(t *testing.T) {
	text := `This is a long looong test sentence,
with some big (biiiiig) words!`

	encodedText, encodedWords := EncodeText([]rune(text))
	if encodedText == text {
		t.Errorf("Encoded text '%s' should be different than provided text '%s'!", encodedText, text)
	}

	if len(encodedWords) != 8 {
		t.Errorf("There should be 8 encoded words! %v", encodedWords)
	}

	decodedText := DecodeText([]rune(encodedText), encodedWords)
	if decodedText != text {
		t.Errorf("Decoded text '%s' should be same as provided text '%s'!", decodedText, text)
	}
}
