package standard

import (
	. "github.com/PaulioRandall/scarlet-go/pkg/token"
)

// SanitiseAll removes redundant tokens, such as comment and whitespace, as well
// as applying formatting to values, e.g trimming off the quotes from string
// literals and templates.
func SanitiseAll(in []Token) (out []Token) {

	var prev Token
	size := len(in)

	for i := 0; i < size; i++ {

		tk := sanitise(prev, in[i])

		if tk != nil {
			out = append(out, tk)
			prev = tk
		}
	}

	return out
}

func sanitise(prev, next Token) Token {

	if isParsableToken(prev, next) {
		return formatToken(next)
	}

	return nil
}

func isParsableToken(prev, next Token) bool {

	if next.Morpheme().Redundant() {
		return false
	}

	if next.Morpheme() == NEWLINE || next.Morpheme() == TERMINATOR {

		if prev == nil {
			return false
		}

		switch prev.Morpheme() {
		// Sometimes the extra terminator or newline is redundant.
		// Removing them makes parsing easier.
		case TERMINATOR,
			DELIMITER,
			BLOCK_OPEN,
			BLOCK_CLOSE,
			MATCH,
			LIST,
			UNDEFINED:

			return false
		}
	}

	return true
}

func formatToken(tk Token) Token {

	switch tk.Morpheme() {
	case NEWLINE:
		// Non-redundant newline tokens are expression and statement terminators
		// in disguise.
		return NewToken(TERMINATOR, tk.Value(), tk.Line(), tk.Col())

	case STRING, TEMPLATE:
		v := tk.Value()
		v = v[1 : len(v)-1] // Remove quotes
		return NewToken(tk.Morpheme(), v, tk.Line(), tk.Col())
	}

	return tk
}