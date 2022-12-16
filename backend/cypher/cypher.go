package cypher

import (
	"cs-be/utilities"
)

const (
	// In a production level I would use my secret keys as a secret argument, or an env var.
	SECRET_KEY_CAESAR         = 10
	SECRET_KEY_MONOALPHABETIC = "AQJGETWLYDHXCBPVAIMUZFOSRN"
	ENCODE                    = 1
	DECODE                    = -1
)

type Text struct {
	Message string `json:"message"`
}

func (t *Text) Caesar(direction int) string {
	rune := utilities.ShiftString(t.Message, direction, SECRET_KEY_CAESAR)
	return string(rune)
}

func (t *Text) Monoalphabetic(direction int) string {
	var finalString string
	for idx, character := range t.Message {
		key := SECRET_KEY_MONOALPHABETIC[idx%len(SECRET_KEY_MONOALPHABETIC)] - 'A' + 1
		output := utilities.ShiftString(string(character), direction, int(key))
		finalString += string(output)
	}
	return finalString
}
