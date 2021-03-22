package weirdtext

import (
	"testing"
)

func TestInvalidPrefix(t *testing.T) {
	input := `
---weir---
Tihs is a lnog lnooog tset setcnnee,
wtih smoe big (biiiiig) wdros!
---weir---
This long looong sentence some test with words`

	test := EncodedText{}
	err := test.FromString(input)
	if err == nil {
		t.Errorf("There should be error as prefix is invalid!")
	}
}

func TestInvalidString(t *testing.T) {
	input := `
---weird---
Tihs is a lnog lnooog tset setcnnee,
wtih smoe big (biiiiig) wdros!
---weird---`
	test := EncodedText{}
	err := test.FromString(input)
	if err == nil {
		t.Errorf("There should be error as string is invalid!")
	}

}
