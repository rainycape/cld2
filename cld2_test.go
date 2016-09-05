package cld2

import "testing"

func TestDetect(t *testing.T) {
	cases := []struct {
		text string
		lang string
	}{
		{"Hello there. Nice to see you. ", "en"},
		{"여보세요", "ko"},
		{"Bonjour, mon ami", "fr"},
	}
	for _, c := range cases {
		actual := Detect(c.text)
		if actual != c.lang {
			t.Errorf("text %#v: want %#v, got %#v", c.text, c.lang, actual)
		}
	}
}
