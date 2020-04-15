package utils

import (
	"log"
	"reflect"
	"testing"
)

func TestGetWordsFromStr(t *testing.T) {
	content := `Prologue

The kids sat leaning over the side of the boat, looking into the water. "A huge pike," Anton said confidently.

"With fins this big?" Pashka asked.

Anton didn't reply. Anka also took a look, but saw only her own reflection.

"Be good to take a swim," said Pashka, plunging his arm into the water up to his elbow. "It's cold," he reported.`

	expected := []string{"Prologue", "The", "kids", "sat", "leaning", "over", "the", "side", "of", "the", "boat", "looking", "into", "the",
		"water", "A", "huge", "pike", "Anton", "said", "confidently", "With", "fins", "this", "big", "Pashka", "asked", "Anton", "didn",
		"t", "reply", "Anka", "also", "took", "a", "look", "but", "saw", "only", "her", "own", "reflection", "Be", "good", "to", "take",
		"a", "swim", "said", "Pashka", "plunging", "his", "arm", "into", "the", "water", "up", "to", "his", "elbow", "It", "s", "cold", "he", "reported"}

	actual := GetWordsFromStr(content)

	if !reflect.DeepEqual(expected, actual) {
		log.Println("Tokens calculate failed")
		t.Fail()
	}
}
