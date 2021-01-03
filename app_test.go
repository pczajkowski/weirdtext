package main

import (
	"testing"
)

type testCase struct {
	text  string
	count int
}

func TestDecodeEncode(t *testing.T) {
	testCases := []testCase{
		{"This is a long looong test sentence,\nwith some big (biiiiig) words!", 8},
		{"Pre-translation generally means applying the TM(s) to one or more files as whole instead of moving segment by segment.", 11},
	}

	for _, item := range testCases {
		encodedText, encodedWords := EncodeText([]rune(item.text))
		if encodedText == item.text {
			t.Errorf("Encoded text '%s' should be different than provided text '%s'!", encodedText, item.text)
		}

		count := len(encodedWords)
		if count != item.count {
			t.Errorf("There should be %d encoded words instead of %d! %v", item.count, count, encodedWords)
		}

		decodedText := DecodeText([]rune(encodedText), encodedWords)
		if decodedText != item.text {
			t.Errorf("Decoded text '%s' should be same as provided text '%s'!", decodedText, item.text)
		}
	}
}
