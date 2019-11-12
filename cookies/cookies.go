package cookies

// NewlineRunes returns the number runes representing the next new line if the
// next token is a newline token else zero is returned. I.e. if the next rune
// is `\n` then `1` is returned, if the the next rune are `\r\n` then `2` is
// returned else 0 is returned.
func NewlineRunes(runes []rune, i int) int {

	if runes[i] == '\n' {
		return 1
	}

	if i+1 < len(runes) {
		if runes[i] == '\r' && runes[i+1] == '\n' {
			return 2
		}
	}

	return 0
}
