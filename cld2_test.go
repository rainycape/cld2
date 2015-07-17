package cld2

import "testing"

type TestData struct {
	ExpectLanguageCode string
	ExpectLanguageName string
	ExpectIsReliable   bool
	Text               string
}

var testData = [...]TestData{
	{"en", "ENGLISH", true, "The quick brown fox jumped over the lazy dog"},
	{"fr", "FRENCH", true, "Le rapide renard brun sauta par dessus le chien paresseux"},
	{"de", "GERMAN", true, "Der schnelle braune Fuchs über den faulen Hund sprang"},
	{"es", "SPANISH", true, "el zorro marrón rápido saltó sobre el perro perezoso"},
	{"mk", "MACEDONIAN", true, "брзо кафеава лисица прескокна мрзливи куче"},
	{"zh", "Chinese", true, "敏捷的棕色狐狸跳过了懒狗，目的也许这语料库文本的宽度足以决定"},
	{"ja", "Japanese", true, "速い茶色のキツネは、怠け者の犬を飛び越えました"},
	{"ko", "Korean", true, "빠른 갈색 여우가 게으른 개를 뛰어 넘었다"},
	{"th", "THAI", true, "สุนัขจิ้งจอกสีน้ำตาลได้อย่างรวดเร็วเพิ่มขึ้นกว่าสุนัขขี้เกียจ"},
	{"ar", "ARABIC", true, "قفز الثعلب البني السريع فوق الكلب الكسول"},
	{"iw", "HEBREW", true, "שועל החום הזריז קפץ מעל הכלב העצלן, לכוון אולי קורפוס זה של טקסט הוא רחב מספיק כדי להחליט"},
	{"un", "Unknown", false, "no"},
}

func TestDetect(t *testing.T) {
	for _, input := range testData {
		actualLanguageCode := Detect(input.Text)
		if actualLanguageCode != input.ExpectLanguageCode {
			t.Errorf("expected `%s`, got `%s` (%s)", input.ExpectLanguageCode, actualLanguageCode, input.Text)
		}
	}
}
