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
		{"新型コロナウィルス感染症拡大防止のため，一部在宅勤務等による接触機会削減を実施しています。電話，FAX，郵便，メールフォーム等でのお問い合わせにつきましては，対応にお時間をいただく場合がありますので，予めご了承ください。また，一般の方の建物内への立入は，当面の間ご遠慮いただいております。みなさまにはご不便をおかけしますが，なにとぞご理解・ご協力のほどよろしくお願い申し上げます。", 10},
		{"Lorem ipsum dolor sit amet, est in aliquip aperiam. Sea te feugiat labores voluptatibus, pri ea affert ubique noluisse! No usu dolorum mentitum, sed eros nihil singulis ad? Has liber putent percipit et, vel te iriure intellegat. Eu dicam placerat has, id eos omittam facilisi, pertinacia constituto te has?", 28},
		{"Λορεμ ιπσθμ δολορ σιτ αμετ, λορεμ λαβοραμθσ ει μελ, cονσετετθρ δεφινιτιονεσ αν εοσ. Εα ηαρθμ cονcλθδατθρqθε περ? Νε ηισ vιvενδθμ μανδαμθσ! Αδ τραcτατοσ δισπθτανδο ιθσ, θτ εστ λαθδεμ μινιμθμ ιντελλεγαμ.", 17},
		{"Szła dzieweczka do laseczka\nDo zielunego, do zielonego, do zielonego.\nNapotkała myśliweczka\nBardzo szwarnego, bardzo szwarnego, bardzo szwarnego.", 14},
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

	decodedText := DecodeText(toDecode)
	if decodedText != expected {
		t.Errorf("Decoded text '%s' should be same as expected text '%s'!", decodedText, expected)
	}
}
