package cypher

const (
	// In a production level I would use my secret keys as a secret argument, or an env var.
	SECRET_KEY_CAESAR = 10
	ENCODE            = -1
	DECODE            = 1
)

type Text struct {
	Message string `json:"message"`
}

func (c *Text) Caesar(direction int) string {
	shift, offset := rune(SECRET_KEY_CAESAR), rune(26)

	// string->rune conversion
	runes := []rune(c.Message)

	for index, char := range runes {
		// Iterate over all runes, and perform substitution
		// wherever possible. If the letter is not in the range
		// [1 .. 25], the offset defined above is added or
		// subtracted.
		switch direction {
		case -1: // encoding
			if char >= 'a'+shift && char <= 'z' ||
				char >= 'A'+shift && char <= 'Z' {
				char = char - shift
			} else if char >= 'a' && char < 'a'+shift ||
				char >= 'A' && char < 'A'+shift {
				char = char - shift + offset
			}
		case +1: // decoding
			if char >= 'a' && char <= 'z'-shift ||
				char >= 'A' && char <= 'Z'-shift {
				char = char + shift
			} else if char > 'z'-shift && char <= 'z' ||
				char > 'Z'-shift && char <= 'Z' {
				char = char + shift - offset
			}
		}

		// Above `if`s handle both upper and lower case ASCII
		// characters; anything else is returned as is (includes
		// numbers, punctuation and space).
		runes[index] = char
	}

	return string(runes)
}

func (c *Text) Monoalphabetic() string {
	return ""
}
