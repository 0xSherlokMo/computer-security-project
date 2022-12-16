package cypher

import (
	"strings"

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
	Key     string `json:"key,omitempty"`
}

func (t *Text) Caesar(direction int) string {
	t.writeKey(string(rune(SECRET_KEY_CAESAR) + 'A' - 1))
	rune := utilities.ShiftString(t.Message, direction, t.offset(0))
	return string(rune)
}

func (t *Text) Monoalphabetic(direction int) string {
	var finalString string
	t.writeKey(SECRET_KEY_MONOALPHABETIC)
	for idx, character := range t.Message {
		output := utilities.ShiftString(string(character), direction, t.offset(idx))
		finalString += string(output)
	}
	return finalString
}

func (t *Text) writeKey(DEFAULT_KEY string) {
	t.Key = DEFAULT_KEY
}

func (t *Text) offset(idx int) int {
	key := strings.ToUpper(t.Key)
	return int(key[idx%len(key)]) - 'A' + 1
}
