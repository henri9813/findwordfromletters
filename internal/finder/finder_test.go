package finder

import "testing"

func TestFindWordsByLetters(t *testing.T) {
	_, err := FindWordsByLetters("kayak")
	if err != nil {
		t.Fatal(err)
	}
}
