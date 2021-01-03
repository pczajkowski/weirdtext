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
		encoded := EncodeText(item.text)
		if encoded.text == item.text {
			t.Errorf("Encoded text '%s' should be different than provided text '%s'!", encoded.text, item.text)
		}

		count := len(encoded.encodedWords)
		if count != item.count {
			t.Errorf("There should be %d encoded words instead of %d! %v", item.count, count, encoded.encodedWords)
		}

		decodedText := DecodeText(encoded.text, encoded.encodedWords)
		if decodedText != item.text {
			t.Errorf("Decoded text '%s' should be same as provided text '%s'!", decodedText, item.text)
		}
	}
}
