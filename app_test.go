package weirdtext

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
		{"Szła dzieweczka do laseczka\nDo zielonego, do zielonego, do zielonego.\nNapotkała myśliweczka\nBardzo szwarnego, bardzo szwarnego, bardzo szwarnego.", 14},
	}

	for _, item := range testCases {
		encoded := EncodeText(item.text)
		if encoded.Text == item.text {
			t.Errorf("Encoded text '%s' should be different than provided text '%s'!", encoded.Text, item.text)
		}

		count := len(encoded.EncodedWords)
		if count != item.count {
			t.Errorf("There should be %d encoded words instead of %d! %v", item.count, count, encoded.EncodedWords)
		}

		decodedText := DecodeText(encoded)
		if decodedText != item.text {
			t.Errorf("Decoded text '%s' should be same as expected text '%s'!", decodedText, item.text)
		}
	}
}

func TestSerializeDeserializeEncodedText(t *testing.T) {
	expected := "This is a long looong test sentence,\nwith some big (biiiiig) words!"

	encoded := EncodeText(expected)
	serialized := encoded.String()

	toDecode := EncodedText{}
	err := toDecode.FromString(serialized)
	if err != nil {
		t.Errorf("Error deserializing encoded text: %s", err)
	}

	decodedText := DecodeText(encoded)
	if decodedText != expected {
		t.Errorf("Decoded text '%s' should be same as expected text '%s'!", decodedText, expected)
	}
}
